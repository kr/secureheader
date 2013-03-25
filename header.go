package secureheader

import (
	"net/http"
)

// HeaderFilter is a Filter that adds a single HTTP header field
// to the response before calling the underlying http.Handler.
type HeaderFilter struct {
	Field, Value string
}

// Filter returns a handler that adds the HTTP header field in f
// to the response, then calls h.ServeHTTP.
func (f HeaderFilter) Filter(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(f.Field, f.Value)
		h.ServeHTTP(w, r)
	})
}
