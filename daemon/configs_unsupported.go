//go:build !linux && !windows

package daemon // import "github.com/Prakhar-Agarwal-byte/moby/daemon"

func configsSupported() bool {
	return false
}
