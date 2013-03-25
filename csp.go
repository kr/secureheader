package secureheader

import (
//"net/http"
//"strings"
)

const (
	firefoxCSPHeader  = "X-Content-Security-Policy"
	webkitCSPHeader   = "X-Webkit-Csp"
	standardCSPHeader = "Content-Security-Policy"
)

const (
	firefoxCSPValue = "options eval-script inline-script; allow https://* data:; frame-src https://* about: javascript:; img-src chrome-extension:"
	webkitCSPValue  = "default-src https: data: 'unsafe-inline' 'unsafe-eval'; frame-src https://* about: javascript:; img-src chrome-extension:"
)

//func NewHandlerCSP() http.Handler {
//	h := new(handlerHeader)
//
//	ua := strings.ToLower(r.Header.Get("User-Agent"))
//	if strings.Contains(ua, "msie") {
//		h.k = standardCSPHeader
//	} else if strings.Contains(ua, "firefox") {
//		h.k = firefoxCSPHeader
//	} else {
//		h.k = webkitCSPHeader
//	}
//
//	if config == nil {
//		if standard {
//			h.v = webkitCSPValue
//		} else {
//			h.v = firefoxCSPValue
//		}
//	} else {
//	}
//
//	if reportOnly {
//		h.k += "-Report-Only"
//	}
//	return h
//}
