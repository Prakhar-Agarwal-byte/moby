//go:build !windows

package snapshotter

import (
	"github.com/containerd/containerd/mount"
	"github.com/Prakhar-Agarwal-byte/moby/daemon/graphdriver"
	"golang.org/x/sys/unix"
)

func checker() graphdriver.Checker {
	return graphdriver.NewDefaultChecker()
}

func unmount(target string) error {
	return mount.Unmount(target, unix.MNT_DETACH)
}
