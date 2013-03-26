// Package secureheader adds some HTTP header fields widely
// considered to improve safety of HTTP requests. These fields
// are documented as follows:
//
//   Strict Transport Security: https://tools.ietf.org/html/rfc6797
//   Frame Options:             https://tools.ietf.org/html/draft-ietf-websec-x-frame-options-00
//   Cross Site Scripting:      http://msdn.microsoft.com/en-us/library/dd565647%28v=vs.85%29.aspx
//   Content Type Options:      http://msdn.microsoft.com/en-us/library/ie/gg622941%28v=vs.85%29.aspx
//
// The easiest way to use this package is to replace nil in your
// http.ListenAndServe call with secureheader.DefaultConfig.
// DefaultConfig is configured with conservative (safer and more
// restrictive) behavior. If you want to customize its behavior,
// assign different values to its fields before calling
// ListenAndServe. See the example code below.
//
// This package was inspired by Twitter's secureheaders Ruby
// library. See https://github.com/twitter/secureheaders.
package secureheader

// TODO(kr): figure out how to add this one:
//   Content Security Policy:   https://dvcs.w3.org/hg/content-security-policy/raw-file/tip/csp-specification.dev.html
// See https://github.com/kr/secureheader/issues/1.

import (
	"net/http"
	"strconv"
	"time"
)

// DefaultConfig wraps the default http handler with some
// conservative (safer and more restrictive) behavior.
var DefaultConfig = &Config{
	HTTPSRedirect: true,

	ContentTypeOptions: true,

	HSTS:                  true,
	HSTSMaxAge:            100 * 24 * time.Hour,
	HSTSIncludeSubdomains: true,

	FrameOptions:      true,
	FrameOpionsPolicy: Deny,

	XSSProtection:      true,
	XSSProtectionBlock: false,

	Next: http.DefaultServeMux,
}

type Config struct {
	// If true, redirects any request with scheme http to the
	// equivalent https URL.
	HTTPSRedirect bool

	// If true, sets X-Content-Type-Options to "nosniff".
	ContentTypeOptions bool

	// If true, sets the HTTP Strict Transport Security header
	// field, which instructs browsers to send future requests
	// over HTTPS, even if the URL uses the unencrypted http
	// scheme.
	HSTS                  bool
	HSTSMaxAge            time.Duration
	HSTSIncludeSubdomains bool

	// If true, adds X-Frame-Options, to control when the request
	// should be displayed inside an HTML frame.
	FrameOptions      bool
	FrameOpionsPolicy FramePolicy

	// If true, sets X-XSS-Protection to "1", optionally with
	// "mode=block". See the official documentation, linked above,
	// for the meaning of these values.
	XSSProtection      bool
	XSSProtectionBlock bool

	Next http.Handler
}

// ServeHTTP sets header fields on w according to the options in
// c, then calls c.Next(w, r). If c.HTTPSRedirect is true and r is
// an unencrypted request, ServeHTTP responds with status 301 and
// does not call c.Next.
func (c *Config) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if c.HTTPSRedirect && r.URL.Scheme != "https" {
		url := *r.URL
		url.Scheme = "https"
		http.Redirect(w, r, url.String(), http.StatusMovedPermanently)
		return
	}
	if c.ContentTypeOptions {
		w.Header().Set("X-Content-Type-Options", "nosniff")
	}
	if c.HSTS {
		v := "max-age=" + strconv.FormatInt(int64(c.HSTSMaxAge/time.Second), 10)
		if c.HSTSIncludeSubdomains {
			v += "; includeSubDomains"
		}
		w.Header().Set("Strict-Transport-Security", v)
	}
	if c.FrameOptions {
		w.Header().Set("X-Frame-Options", string(c.FrameOpionsPolicy))
	}
	if c.XSSProtection {
		v := "1"
		if c.XSSProtectionBlock {
			v += "; mode=block"
		}
		w.Header().Set("X-XSS-Protection", v)
	}
	c.Next.ServeHTTP(w, r)
}

// FramePolicy tells the browser under what circumstances to allow
// the response to be displayed inside an HTML frame. There are
// three options:
//
//   Deny            do not permit display in a frame
//   SameOrigin      permit display in a frame from the same origin
//   AllowFrom(url)  permit display in a frame from the given url
type FramePolicy string

const (
	Deny       FramePolicy = "DENY"
	SameOrigin FramePolicy = "SAMEORIGIN"
)

// AllowFrom returns a FramePolicy specifying that each requested
// resource can be included in a frame from only the given url.
func AllowFrom(url string) FramePolicy {
	return FramePolicy("ALLOW-FROM: " + url)
}
