package daemon // import "github.com/Prakhar-Agarwal-byte/moby/daemon"

import (
	"context"

	"github.com/Prakhar-Agarwal-byte/moby/api/types/registry"
	"github.com/Prakhar-Agarwal-byte/moby/dockerversion"
)

// AuthenticateToRegistry checks the validity of credentials in authConfig
func (daemon *Daemon) AuthenticateToRegistry(ctx context.Context, authConfig *registry.AuthConfig) (string, string, error) {
	return daemon.registryService.Auth(ctx, authConfig, dockerversion.DockerUserAgent(ctx))
}
