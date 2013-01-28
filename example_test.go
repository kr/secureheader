package secureheaders_test

import (
	"github.com/kr/secureheaders"
	"net/http"
	"time"
)

func Example() {
	http.Handle("/", http.FileServer(http.Dir("/tmp")))
	http.ListenAndServe(":80", secureheaders.DefaultHandler)
}

func Example_custom() {
	http.Handle("/", http.FileServer(http.Dir("/tmp")))

	// Permit unencrypted HTTP requests to subdomains.
	// Permit pages to be contained in frames (omit FrameOptions).
	f := secureheaders.Compose(
		secureheaders.HTTPSRedirect(),
		secureheaders.ContentTypeOptions(),
		secureheaders.HSTS(100*24*time.Hour, false),
		secureheaders.XSSProtection(true, false),
	)
	http.ListenAndServe(":80", f.Filter(http.DefaultServeMux))
}

func ExampleFrameOptionsAllowFrom() {
	http.Handle("/", http.FileServer(http.Dir("/tmp")))
	f := secureheaders.FrameOptionsAllowFrom("https://example.com/")
	http.ListenAndServe(":80", f.Filter(http.DefaultServeMux))
}
