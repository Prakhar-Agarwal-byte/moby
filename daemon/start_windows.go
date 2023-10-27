package daemon // import "github.com/Prakhar-Agarwal-byte/moby/daemon"

import (
	"github.com/Microsoft/hcsshim/cmd/containerd-shim-runhcs-v1/options"
	"github.com/Prakhar-Agarwal-byte/moby/container"
	"github.com/Prakhar-Agarwal-byte/moby/pkg/system"
)

func (daemon *Daemon) getLibcontainerdCreateOptions(*configStore, *container.Container) (string, interface{}, error) {
	if system.ContainerdRuntimeSupported() {
		opts := &options.Options{}
		return "io.containerd.runhcs.v1", opts, nil
	}
	return "", nil, nil
}
