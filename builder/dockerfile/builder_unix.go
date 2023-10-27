//go:build !windows

package dockerfile // import "github.com/Prakhar-Agarwal-byte/moby/builder/dockerfile"

func defaultShellForOS(os string) []string {
	return []string{"/bin/sh", "-c"}
}
