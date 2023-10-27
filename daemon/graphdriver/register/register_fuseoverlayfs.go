//go:build !exclude_graphdriver_fuseoverlayfs && linux

package register // import "github.com/Prakhar-Agarwal-byte/moby/daemon/graphdriver/register"

import (
	// register the fuse-overlayfs graphdriver
	_ "github.com/Prakhar-Agarwal-byte/moby/daemon/graphdriver/fuse-overlayfs"
)
