package secureheader

// FrameOptions returns a filter that adds an X-Frame-Options
// header field with policy p.
func frameOptions(p string) Filter {
	return &HeaderFilter{"X-Frame-Options", p}
}

// FrameOptions returns a filter that adds an X-Frame-Options
// header field with a policy of either SAMEORIGIN or DENY.
func FrameOptions(deny bool) Filter {
	if deny {
		return frameOptions("DENY")
	}
	return frameOptions("SAMEORIGIN")
}

// FrameOptionsAllowFrom returns a filter that adds an
// X-Frame-Options header field specifying each requested resource
// can be included in a frame from only the given url.
func FrameOptionsAllowFrom(url string) Filter {
	return frameOptions("ALLOW-FROM: " + url)
}
