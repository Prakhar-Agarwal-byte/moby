//go:build !linux && !freebsd

package zfs // import "github.com/Prakhar-Agarwal-byte/moby/daemon/graphdriver/zfs"

func checkRootdirFs(rootdir string) error {
	return nil
}

func getMountpoint(id string) string {
	return id
}
