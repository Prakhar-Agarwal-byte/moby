package daemon // import "github.com/Prakhar-Agarwal-byte/moby/daemon"

import (
	"context"

	"github.com/Prakhar-Agarwal-byte/moby/api/types"
	"github.com/Prakhar-Agarwal-byte/moby/api/types/backend"
	"github.com/Prakhar-Agarwal-byte/moby/container"
)

// This sets platform-specific fields
func setPlatformSpecificContainerFields(container *container.Container, contJSONBase *types.ContainerJSONBase) *types.ContainerJSONBase {
	return contJSONBase
}

// containerInspectPre120 get containers for pre 1.20 APIs.
func (daemon *Daemon) containerInspectPre120(ctx context.Context, name string) (*types.ContainerJSON, error) {
	return daemon.ContainerInspectCurrent(ctx, name, false)
}

func inspectExecProcessConfig(e *container.ExecConfig) *backend.ExecProcessConfig {
	return &backend.ExecProcessConfig{
		Tty:        e.Tty,
		Entrypoint: e.Entrypoint,
		Arguments:  e.Args,
	}
}
