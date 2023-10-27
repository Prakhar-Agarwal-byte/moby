//go:build linux || freebsd || darwin

package builtin

import (
	"github.com/Prakhar-Agarwal-byte/moby/libnetwork/ipamapi"
)

// Register registers the built-in ipam service with libnetwork.
func Register(r ipamapi.Registerer) error {
	return registerBuiltin(r)
}
