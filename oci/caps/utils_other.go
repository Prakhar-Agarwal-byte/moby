//go:build !linux

package caps // import "github.com/Prakhar-Agarwal-byte/moby/oci/caps"

func initCaps() {
	// no capabilities on Windows
}
