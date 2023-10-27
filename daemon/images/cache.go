package images // import "github.com/Prakhar-Agarwal-byte/moby/daemon/images"

import (
	"context"

	"github.com/containerd/log"
	imagetypes "github.com/Prakhar-Agarwal-byte/moby/api/types/image"
	"github.com/Prakhar-Agarwal-byte/moby/builder"
	"github.com/Prakhar-Agarwal-byte/moby/image/cache"
	"github.com/pkg/errors"
)

// MakeImageCache creates a stateful image cache.
func (i *ImageService) MakeImageCache(ctx context.Context, sourceRefs []string) (builder.ImageCache, error) {
	if len(sourceRefs) == 0 {
		return cache.NewLocal(i.imageStore), nil
	}

	cache := cache.New(i.imageStore)

	for _, ref := range sourceRefs {
		img, err := i.GetImage(ctx, ref, imagetypes.GetImageOpts{})
		if err != nil {
			if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
				return nil, err
			}
			log.G(ctx).Warnf("Could not look up %s for cache resolution, skipping: %+v", ref, err)
			continue
		}
		cache.Populate(img)
	}

	return cache, nil
}
