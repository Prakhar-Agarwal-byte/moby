//go:build !exclude_graphdriver_btrfs && linux

package register // import "github.com/Prakhar-Agarwal-byte/moby/daemon/graphdriver/register"

import (
	// register the btrfs graphdriver
	_ "github.com/Prakhar-Agarwal-byte/moby/daemon/graphdriver/btrfs"
)
