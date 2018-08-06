package secureheader_test

import (
	"net/http"

	"github.com/kr/secureheader"
)

func Example() {
	http.Handle("/", http.FileServer(http.Dir("/tmp")))
	http.ListenAndServe(":80", secureheader.Handler(nil))
}

func Example_custom() {
	http.Handle("/", http.FileServer(http.Dir("/tmp")))
	h := secureheader.Handler(http.DefaultServeMux)
	h.HSTSIncludeSubdomains = false
	h.FrameOptions = false
	http.ListenAndServe(":80", h)
}

func Example_Content_Security_Policy() {
	http.Handle("/", http.FileServer(http.Dir("/tmp")))
	h := secureheader.Handler(http.DefaultServeMux)
	h.CSP = true
	h.CSPBody = "default-src 'self' ; img-src 'self' data: ; style-src 'self'"
	h.CSPReportURI = "https://example.com/csp-reports"
	http.ListenAndServe(":80", h)

}
