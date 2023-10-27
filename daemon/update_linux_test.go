package daemon // import "github.com/Prakhar-Agarwal-byte/moby/daemon"

import (
	"testing"

	"github.com/Prakhar-Agarwal-byte/moby/api/types/container"
)

func TestToContainerdResources_Defaults(t *testing.T) {
	checkResourcesAreUnset(t, toContainerdResources(container.Resources{}))
}
