package daemon // import "github.com/Prakhar-Agarwal-byte/moby/daemon"

import (
	"github.com/Prakhar-Agarwal-byte/moby/api/types/system"
	"github.com/Prakhar-Agarwal-byte/moby/dockerversion"
)

func (daemon *Daemon) fillLicense(v *system.Info) {
	v.ProductLicense = dockerversion.DefaultProductLicense
}
