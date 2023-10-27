package distribution // import "github.com/Prakhar-Agarwal-byte/moby/api/server/router/distribution"

import (
	"context"

	"github.com/distribution/reference"
	"github.com/docker/distribution"
	"github.com/Prakhar-Agarwal-byte/moby/api/types/registry"
)

// Backend is all the methods that need to be implemented
// to provide image specific functionality.
type Backend interface {
	GetRepository(context.Context, reference.Named, *registry.AuthConfig) (distribution.Repository, error)
}
