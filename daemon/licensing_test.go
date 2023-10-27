package daemon // import "github.com/Prakhar-Agarwal-byte/moby/daemon"

import (
	"testing"

	"github.com/Prakhar-Agarwal-byte/moby/api/types/system"
	"github.com/Prakhar-Agarwal-byte/moby/dockerversion"
	"gotest.tools/v3/assert"
)

func TestFillLicense(t *testing.T) {
	v := &system.Info{}
	d := &Daemon{
		root: "/var/lib/docker/",
	}
	d.fillLicense(v)
	assert.Assert(t, v.ProductLicense == dockerversion.DefaultProductLicense)
}
