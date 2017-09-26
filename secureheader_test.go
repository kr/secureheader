package secureheader

import (
	"testing"
)

func TestDefaultUseForwardedProto(t *testing.T) {
	if defaultUseForwardedProto {
		t.Fatal("defaultUseForwardedProto = true want false")
	}
}

func TestDefaultConfigHTTPSRedirect(t *testing.T) {
	if !DefaultConfig.HTTPSRedirect {
		t.Fatal("HTTPSRedirect = false want true")
	}
}

func TestDefaultHSTSIncludeSubdomains(t *testing.T) {
	if !DefaultConfig.HSTSIncludeSubdomains {
		t.Fatal("HSTSIncludeSubdomains = false want true")
	}
}

func TestDefaultHSTSPreload(t *testing.T) {
	if DefaultConfig.HSTSPreload {
		t.Fatal("HSTSPreload = true want false")
	}
}
