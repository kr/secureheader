package secureheader

import (
	"fmt"
	"time"
)

// HSTS returns a filter that sets the HTTP Strict Transport
// Security header field, which instructs browsers to send future
// requests over HTTPS, even if the URL uses the unencrypted http
// scheme.
func HSTS(maxAge time.Duration, includeSubdomains bool) Filter {
	f := new(HeaderFilter)
	f.Field = "Strict-Transport-Security"
	f.Value = fmt.Sprintf("max-age=%d", maxAge/time.Second)
	if includeSubdomains {
		f.Value += "; includeSubDomains"
	}
	return f
}
