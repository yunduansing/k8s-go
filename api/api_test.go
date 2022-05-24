package api

import (
	"os"
	"testing"
)

func TestReq(t *testing.T) {
	os.Setenv("KUBERNETES_SERVICE_HOST", "localhost")
	os.Setenv("KUBERNETES_SERVICE_PORT", "30302")
	Req()
}
