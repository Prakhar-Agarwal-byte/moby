package vfs // import "github.com/Prakhar-Agarwal-byte/moby/daemon/graphdriver/vfs"

import "github.com/Prakhar-Agarwal-byte/moby/daemon/graphdriver/copy"

func dirCopy(srcDir, dstDir string) error {
	return copy.DirCopy(srcDir, dstDir, copy.Content, false)
}
