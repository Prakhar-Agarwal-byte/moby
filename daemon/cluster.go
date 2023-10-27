package daemon // import "github.com/Prakhar-Agarwal-byte/moby/daemon"

import (
	apitypes "github.com/Prakhar-Agarwal-byte/moby/api/types"
	"github.com/Prakhar-Agarwal-byte/moby/api/types/filters"
	lncluster "github.com/Prakhar-Agarwal-byte/moby/libnetwork/cluster"
)

// Cluster is the interface for github.com/Prakhar-Agarwal-byte/moby/daemon/cluster.(*Cluster).
type Cluster interface {
	ClusterStatus
	NetworkManager
	SendClusterEvent(event lncluster.ConfigEventType)
}

// ClusterStatus interface provides information about the Swarm status of the Cluster
type ClusterStatus interface {
	IsAgent() bool
	IsManager() bool
}

// NetworkManager provides methods to manage networks
type NetworkManager interface {
	GetNetwork(input string) (apitypes.NetworkResource, error)
	GetNetworks(filters.Args) ([]apitypes.NetworkResource, error)
	RemoveNetwork(input string) error
}
