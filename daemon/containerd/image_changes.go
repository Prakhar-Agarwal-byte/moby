package containerd

import (
	"context"

	"github.com/Prakhar-Agarwal-byte/moby/container"
	"github.com/Prakhar-Agarwal-byte/moby/pkg/archive"
	"github.com/Prakhar-Agarwal-byte/moby/pkg/stringid"
	"github.com/containerd/containerd/mount"
	"github.com/containerd/log"
)

func (i *ImageService) Changes(ctx context.Context, container *container.Container) ([]archive.Change, error) {
	snapshotter := i.client.SnapshotService(container.Driver)
	info, err := snapshotter.Stat(ctx, container.ID)
	if err != nil {
		return nil, err
	}

	id := stringid.GenerateRandomID()
	parentViewKey := container.ID + "-parent-view-" + id
	imageMounts, _ := snapshotter.View(ctx, parentViewKey, info.Parent)

	defer func() {
		if err := snapshotter.Remove(ctx, parentViewKey); err != nil {
			log.G(ctx).WithError(err).Warn("error removing the parent view snapshot")
		}
	}()

	var changes []archive.Change
	err = i.PerformWithBaseFS(ctx, container, func(containerRoot string) error {
		return mount.WithReadonlyTempMount(ctx, imageMounts, func(imageRoot string) error {
			changes, err = archive.ChangesDirs(containerRoot, imageRoot)
			return err
		})
	})

	return changes, err
}
