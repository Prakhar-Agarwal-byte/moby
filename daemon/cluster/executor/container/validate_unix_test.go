//go:build !windows

package container // import "github.com/Prakhar-Agarwal-byte/moby/daemon/cluster/executor/container"

const (
	testAbsPath        = "/foo"
	testAbsNonExistent = "/some-non-existing-host-path/"
)
