//go:build !linux && !windows

package daemon // import "github.com/Prakhar-Agarwal-byte/moby/daemon"

func secretsSupported() bool {
	return false
}
