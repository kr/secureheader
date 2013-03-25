package secureheader

// ContentTypeOptions returns a filter that sets header field
// X-Content-Type-Options to "nosniff".
func ContentTypeOptions() Filter {
	return &HeaderFilter{"X-Content-Type-Options", "nosniff"}
}
