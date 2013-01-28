// Package secureheaders adds some HTTP headers widely considered
// to improve safety of HTTP requests. These headers are
// documented as follows:
//
//   Strict Transport Security: https://tools.ietf.org/html/rfc6797
//   Frame Options:             https://tools.ietf.org/html/draft-ietf-websec-x-frame-options-00
//   Cross Site Scripting:      http://msdn.microsoft.com/en-us/library/dd565647%28v=vs.85%29.aspx
//   Content Type Options:      http://msdn.microsoft.com/en-us/library/ie/gg622941%28v=vs.85%29.aspx
//
// The easiest way to use this package is to replace nil in your
// http.ListenAndServe call with secureheaders.DefaultHandler.
// DefaultHandler is initialized to a chain of handlers with
// conservative (safer and more restrictive) behavior. If you want
// to customize its behavior (for example, omitting one of the
// handlers in the chain or changing some of the handlers'
// parameters), copy the expression that was used to set
// DefaultHandler, edit it, and use it to make your own chain of
// handlers (as in the "Custom" example).
//
// This package was inspired by Twitter's secureheaders Ruby
// library. See https://github.com/twitter/secureheaders.
package secureheaders

// TODO(kr): figure out how to add this one:
//   Content Security Policy:   https://dvcs.w3.org/hg/content-security-policy/raw-file/tip/csp-specification.dev.html

import (
	"net/http"
	"time"
)

// DefaultFilter is a chain of filters with conservative (safer
// and more restrictive) behavior.
var DefaultFilter Filter = Compose(
	HTTPSRedirect(),
	ContentTypeOptions(),
	HSTS(100*24*time.Hour, true),
	FrameOptions(true),
	XSSProtection(true, false),
)

// DefaultHandler wraps the default http handler with some
// conservative (safer and more restrictive) behavior.
var DefaultHandler = DefaultFilter.Filter(http.DefaultServeMux)
