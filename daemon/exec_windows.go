package daemon // import "github.com/Prakhar-Agarwal-byte/moby/daemon"

import (
	"context"

	"github.com/Prakhar-Agarwal-byte/moby/container"
	"github.com/Prakhar-Agarwal-byte/moby/daemon/config"
	specs "github.com/opencontainers/runtime-spec/specs-go"
)

func (daemon *Daemon) execSetPlatformOpt(ctx context.Context, daemonCfg *config.Config, ec *container.ExecConfig, p *specs.Process) error {
	if ec.Container.OS == "windows" {
		p.User.Username = ec.User
	}
	return nil
}
