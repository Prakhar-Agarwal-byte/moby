//go:build linux || freebsd || darwin || openbsd

package layer // import "github.com/Prakhar-Agarwal-byte/moby/layer"

import "github.com/Prakhar-Agarwal-byte/moby/pkg/stringid"

func (ls *layerStore) mountID(name string) string {
	return stringid.GenerateRandomID()
}
