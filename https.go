package secureheaders

import (
	"net/http"
)

// HTTPSRedirect returns a simple filter that redirects any
// request with scheme http to the equivalent https URL.
func HTTPSRedirect() Filter {
	return FilterFunc(filterHTTPSRedirect)
}

func filterHTTPSRedirect(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Scheme != "https" {
			url := *r.URL
			url.Scheme = "https"
			http.Redirect(w, r, url.String(), http.StatusMovedPermanently)
			return
		}
		h.ServeHTTP(w, r)
	})
}
