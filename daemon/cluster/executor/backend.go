package executor // import "github.com/Prakhar-Agarwal-byte/moby/daemon/cluster/executor"

import (
	"context"
	"io"
	"time"

	"github.com/Prakhar-Agarwal-byte/moby/api/types"
	"github.com/Prakhar-Agarwal-byte/moby/api/types/backend"
	"github.com/Prakhar-Agarwal-byte/moby/api/types/container"
	"github.com/Prakhar-Agarwal-byte/moby/api/types/events"
	"github.com/Prakhar-Agarwal-byte/moby/api/types/filters"
	opts "github.com/Prakhar-Agarwal-byte/moby/api/types/image"
	"github.com/Prakhar-Agarwal-byte/moby/api/types/network"
	"github.com/Prakhar-Agarwal-byte/moby/api/types/registry"
	"github.com/Prakhar-Agarwal-byte/moby/api/types/swarm"
	"github.com/Prakhar-Agarwal-byte/moby/api/types/system"
	"github.com/Prakhar-Agarwal-byte/moby/api/types/volume"
	containerpkg "github.com/Prakhar-Agarwal-byte/moby/container"
	clustertypes "github.com/Prakhar-Agarwal-byte/moby/daemon/cluster/provider"
	networkSettings "github.com/Prakhar-Agarwal-byte/moby/daemon/network"
	"github.com/Prakhar-Agarwal-byte/moby/image"
	"github.com/Prakhar-Agarwal-byte/moby/libnetwork"
	"github.com/Prakhar-Agarwal-byte/moby/libnetwork/cluster"
	networktypes "github.com/Prakhar-Agarwal-byte/moby/libnetwork/types"
	"github.com/Prakhar-Agarwal-byte/moby/plugin"
	volumeopts "github.com/Prakhar-Agarwal-byte/moby/volume/service/opts"
	"github.com/distribution/reference"
	"github.com/docker/distribution"
	"github.com/moby/swarmkit/v2/agent/exec"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
)

// Backend defines the executor component for a swarm agent.
type Backend interface {
	CreateManagedNetwork(clustertypes.NetworkCreateRequest) error
	DeleteManagedNetwork(networkID string) error
	FindNetwork(idName string) (*libnetwork.Network, error)
	SetupIngress(clustertypes.NetworkCreateRequest, string) (<-chan struct{}, error)
	ReleaseIngress() (<-chan struct{}, error)
	CreateManagedContainer(ctx context.Context, config types.ContainerCreateConfig) (container.CreateResponse, error)
	ContainerStart(ctx context.Context, name string, hostConfig *container.HostConfig, checkpoint string, checkpointDir string) error
	ContainerStop(ctx context.Context, name string, config container.StopOptions) error
	ContainerLogs(ctx context.Context, name string, config *container.LogsOptions) (msgs <-chan *backend.LogMessage, tty bool, err error)
	ConnectContainerToNetwork(containerName, networkName string, endpointConfig *network.EndpointSettings) error
	ActivateContainerServiceBinding(containerName string) error
	DeactivateContainerServiceBinding(containerName string) error
	UpdateContainerServiceConfig(containerName string, serviceConfig *clustertypes.ServiceConfig) error
	ContainerInspectCurrent(ctx context.Context, name string, size bool) (*types.ContainerJSON, error)
	ContainerWait(ctx context.Context, name string, condition containerpkg.WaitCondition) (<-chan containerpkg.StateStatus, error)
	ContainerRm(name string, config *types.ContainerRmConfig) error
	ContainerKill(name string, sig string) error
	SetContainerDependencyStore(name string, store exec.DependencyGetter) error
	SetContainerSecretReferences(name string, refs []*swarm.SecretReference) error
	SetContainerConfigReferences(name string, refs []*swarm.ConfigReference) error
	SystemInfo() *system.Info
	Containers(ctx context.Context, config *container.ListOptions) ([]*types.Container, error)
	SetNetworkBootstrapKeys([]*networktypes.EncryptionKey) error
	DaemonJoinsCluster(provider cluster.Provider)
	DaemonLeavesCluster()
	IsSwarmCompatible() error
	SubscribeToEvents(since, until time.Time, filter filters.Args) ([]events.Message, chan interface{})
	UnsubscribeFromEvents(listener chan interface{})
	UpdateAttachment(string, string, string, *network.NetworkingConfig) error
	WaitForDetachment(context.Context, string, string, string, string) error
	PluginManager() *plugin.Manager
	PluginGetter() *plugin.Store
	GetAttachmentStore() *networkSettings.AttachmentStore
	HasExperimental() bool
}

// VolumeBackend is used by an executor to perform volume operations
type VolumeBackend interface {
	Create(ctx context.Context, name, driverName string, opts ...volumeopts.CreateOption) (*volume.Volume, error)
}

// ImageBackend is used by an executor to perform image operations
type ImageBackend interface {
	PullImage(ctx context.Context, ref reference.Named, platform *ocispec.Platform, metaHeaders map[string][]string, authConfig *registry.AuthConfig, outStream io.Writer) error
	GetRepository(context.Context, reference.Named, *registry.AuthConfig) (distribution.Repository, error)
	GetImage(ctx context.Context, refOrID string, options opts.GetImageOpts) (*image.Image, error)
}
