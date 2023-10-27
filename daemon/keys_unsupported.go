//go:build !linux

package daemon // import "github.com/Prakhar-Agarwal-byte/moby/daemon"

// modifyRootKeyLimit is a noop on unsupported platforms.
func modifyRootKeyLimit() error {
	return nil
}
