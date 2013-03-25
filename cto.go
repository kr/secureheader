package secureheaders

// ContentTypeOptions returns a filter that sets header field
// X-Content-Type-Options to "nosniff".
func ContentTypeOptions() Filter {
	return &HeaderFilter{
		Field: "X-Content-Type-Options",
		Value: "nosniff",
	}
}
