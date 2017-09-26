package secureheader_test

import (
	"github.com/kr/secureheader"
	"net/http"
)

func Example() {
	http.Handle("/", http.FileServer(http.Dir("/tmp")))
	http.ListenAndServe(":80", secureheader.DefaultConfig)
}

func Example_custom() {
	http.Handle("/", http.FileServer(http.Dir("/tmp")))
	secureheader.DefaultConfig.HSTSIncludeSubdomains = false
	secureheader.DefaultConfig.FrameOptions = false
	http.ListenAndServe(":80", secureheader.DefaultConfig)
}

func Example_Content_Security_Policy() {
	http.Handle("/", http.FileServer(http.Dir("/tmp")))
	secureheader.DefaultConfig.CSP = true
	secureheader.DefaultConfig.CSPBody = "default-src 'self' ; img-src 'self' data: ; style-src 'self'"
	secureheader.DefaultConfig.CSPReportURI = "https://example.com/csp-reports"
	http.ListenAndServe(":80", secureheader.DefaultConfig)

}
