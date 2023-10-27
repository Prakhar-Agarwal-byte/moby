//go:build !windows

package containerfs // import "github.com/Prakhar-Agarwal-byte/moby/pkg/containerfs"

import "path/filepath"

// CleanScopedPath preappends a to combine with a mnt path.
func CleanScopedPath(path string) string {
	return filepath.Join(string(filepath.Separator), path)
}
