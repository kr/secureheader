// +build heroku

package secureheader

func useForwardedProto() bool {
	return true
}
