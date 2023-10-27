//go:build !linux

package daemon // import "github.com/Prakhar-Agarwal-byte/moby/daemon"

func ensureDefaultAppArmorProfile() error {
	return nil
}

// DefaultApparmorProfile returns an empty string.
func DefaultApparmorProfile() string {
	return ""
}
