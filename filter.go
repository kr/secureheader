package secureheader

import (
	"net/http"
)

// A Filter takes an http.Handler and returns a new http.Handler.
// Typically the returned handler will alter the request or
// response in some way, then call h.ServeHTTP.
type Filter interface {
	Filter(h http.Handler) http.Handler
}

// Represents a chain of handlers, left to right, ending at h.
// If len is 0, this is the identity filter.
type composition []Filter

func (c composition) Filter(h http.Handler) http.Handler {
	if len(c) < 1 {
		return h
	}
	return c[0].Filter(c[1:].Filter(h))
}

// Returns the composition of the given filters f.
func Compose(f ...Filter) Filter {
	return composition(f)
}

// The FilterFunc type is an adapter to allow the use of ordinary
// functions as HTTP handler filters. If f is a function with the
// appropriate signature, FilterFunc(f) is a Filter object that
// calls f.
type FilterFunc func(h http.Handler) http.Handler

// Filter calls f(h).
func (f FilterFunc) Filter(h http.Handler) http.Handler {
	return f(h)
}
