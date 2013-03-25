package secureheader_test

import (
	"github.com/kr/secureheader"
	"net/http"
	"time"
)

func Example() {
	http.Handle("/", http.FileServer(http.Dir("/tmp")))
	http.ListenAndServe(":80", secureheader.DefaultHandler)
}

func Example_custom() {
	http.Handle("/", http.FileServer(http.Dir("/tmp")))

	// Differences from DefaultFilter:
	// - permit unencrypted HTTP requests to subdomains in HSTS
	// - permit pages to be contained in frames (omit FrameOptions)
	f := secureheader.Compose(
		secureheader.HTTPSRedirect(),
		secureheader.ContentTypeOptions(),
		secureheader.HSTS(100*24*time.Hour, false),
		secureheader.XSSProtection(true, false),
	)
	http.ListenAndServe(":80", f.Filter(http.DefaultServeMux))
}

func ExampleFrameOptionsAllowFrom() {
	http.Handle("/", http.FileServer(http.Dir("/tmp")))
	f := secureheader.FrameOptionsAllowFrom("https://example.com/")
	http.ListenAndServe(":80", f.Filter(http.DefaultServeMux))
}
