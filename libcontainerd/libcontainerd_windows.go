package libcontainerd // import "github.com/Prakhar-Agarwal-byte/moby/libcontainerd"

import (
	"context"

	"github.com/Prakhar-Agarwal-byte/moby/libcontainerd/local"
	"github.com/Prakhar-Agarwal-byte/moby/libcontainerd/remote"
	libcontainerdtypes "github.com/Prakhar-Agarwal-byte/moby/libcontainerd/types"
	"github.com/Prakhar-Agarwal-byte/moby/pkg/system"
	"github.com/containerd/containerd"
)

// NewClient creates a new libcontainerd client from a containerd client
func NewClient(ctx context.Context, cli *containerd.Client, stateDir, ns string, b libcontainerdtypes.Backend) (libcontainerdtypes.Client, error) {
	if !system.ContainerdRuntimeSupported() {
		return local.NewClient(ctx, cli, stateDir, ns, b)
	}
	return remote.NewClient(ctx, cli, stateDir, ns, b)
}
