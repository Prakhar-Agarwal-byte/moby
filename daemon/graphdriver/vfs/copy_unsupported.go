//go:build !linux

package vfs // import "github.com/Prakhar-Agarwal-byte/moby/daemon/graphdriver/vfs"

import (
	"github.com/Prakhar-Agarwal-byte/moby/pkg/chrootarchive"
	"github.com/Prakhar-Agarwal-byte/moby/pkg/idtools"
)

func dirCopy(srcDir, dstDir string) error {
	return chrootarchive.NewArchiver(idtools.IdentityMapping{}).CopyWithTar(srcDir, dstDir)
}
