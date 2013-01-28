package secureheaders

// XSSProtection returns a filter that sets header field
// X-XSS-Protection to "1" or "0", optionally with "mode=block".
// See the official documentation, linked above, for the meaning
// of these values.
func XSSProtection(on, block bool) Filter {
	f := new(HeaderFilter)
	f.Field = "X-XSS-Protection"
	if on {
		f.Value =  "1"
	} else {
		f.Value =  "0"
	}
	if block {
		f.Value += "; mode=block"
	}
	return f
}
