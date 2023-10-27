package volume // import "github.com/Prakhar-Agarwal-byte/moby/api/server/router/volume"

import (
	"context"

	"github.com/Prakhar-Agarwal-byte/moby/volume/service/opts"
	// TODO return types need to be refactored into pkg
	"github.com/Prakhar-Agarwal-byte/moby/api/types"
	"github.com/Prakhar-Agarwal-byte/moby/api/types/filters"
	"github.com/Prakhar-Agarwal-byte/moby/api/types/volume"
)

// Backend is the methods that need to be implemented to provide
// volume specific functionality
type Backend interface {
	List(ctx context.Context, filter filters.Args) ([]*volume.Volume, []string, error)
	Get(ctx context.Context, name string, opts ...opts.GetOption) (*volume.Volume, error)
	Create(ctx context.Context, name, driverName string, opts ...opts.CreateOption) (*volume.Volume, error)
	Remove(ctx context.Context, name string, opts ...opts.RemoveOption) error
	Prune(ctx context.Context, pruneFilters filters.Args) (*types.VolumesPruneReport, error)
}

// ClusterBackend is the backend used for Swarm Cluster Volumes. Regular
// volumes go through the volume service, but to avoid across-dependency
// between the cluster package and the volume package, we simply provide two
// backends here.
type ClusterBackend interface {
	GetVolume(nameOrID string) (volume.Volume, error)
	GetVolumes(options volume.ListOptions) ([]*volume.Volume, error)
	CreateVolume(volume volume.CreateOptions) (*volume.Volume, error)
	RemoveVolume(nameOrID string, force bool) error
	UpdateVolume(nameOrID string, version uint64, volume volume.UpdateOptions) error
	IsManager() bool
}
