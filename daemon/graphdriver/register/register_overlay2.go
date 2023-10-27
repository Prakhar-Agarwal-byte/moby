//go:build !exclude_graphdriver_overlay2 && linux

package register // import "github.com/Prakhar-Agarwal-byte/moby/daemon/graphdriver/register"

import (
	// register the overlay2 graphdriver
	_ "github.com/Prakhar-Agarwal-byte/moby/daemon/graphdriver/overlay2"
)
