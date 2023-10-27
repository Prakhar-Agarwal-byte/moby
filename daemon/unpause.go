package daemon // import "github.com/Prakhar-Agarwal-byte/moby/daemon"

import (
	"context"
	"fmt"

	"github.com/Prakhar-Agarwal-byte/moby/api/types/events"
	"github.com/Prakhar-Agarwal-byte/moby/container"
	"github.com/containerd/log"
)

// ContainerUnpause unpauses a container
func (daemon *Daemon) ContainerUnpause(name string) error {
	ctr, err := daemon.GetContainer(name)
	if err != nil {
		return err
	}
	return daemon.containerUnpause(ctr)
}

// containerUnpause resumes the container execution after the container is paused.
func (daemon *Daemon) containerUnpause(ctr *container.Container) error {
	ctr.Lock()
	defer ctr.Unlock()

	// We cannot unpause the container which is not paused
	if !ctr.Paused {
		return fmt.Errorf("Container %s is not paused", ctr.ID)
	}
	tsk, err := ctr.GetRunningTask()
	if err != nil {
		return err
	}

	if err := tsk.Resume(context.Background()); err != nil {
		return fmt.Errorf("Cannot unpause container %s: %s", ctr.ID, err)
	}

	ctr.Paused = false
	daemon.setStateCounter(ctr)
	daemon.updateHealthMonitor(ctr)
	daemon.LogContainerEvent(ctr, events.ActionUnPause)

	if err := ctr.CheckpointTo(daemon.containersReplica); err != nil {
		log.G(context.TODO()).WithError(err).Warn("could not save container to disk")
	}

	return nil
}
