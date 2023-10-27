//go:build !windows

package archive // import "github.com/Prakhar-Agarwal-byte/moby/pkg/archive"

import (
	"path/filepath"
)

func normalizePath(path string) string {
	return filepath.ToSlash(path)
}
