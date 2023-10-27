package daemon // import "github.com/Prakhar-Agarwal-byte/moby/daemon"

import (
	"context"

	swarmtypes "github.com/Prakhar-Agarwal-byte/moby/api/types/swarm"
	"github.com/containerd/log"
)

// SetContainerConfigReferences sets the container config references needed
func (daemon *Daemon) SetContainerConfigReferences(name string, refs []*swarmtypes.ConfigReference) error {
	if !configsSupported() && len(refs) > 0 {
		log.G(context.TODO()).Warn("configs are not supported on this platform")
		return nil
	}

	c, err := daemon.GetContainer(name)
	if err != nil {
		return err
	}
	c.ConfigReferences = append(c.ConfigReferences, refs...)
	return nil
}
