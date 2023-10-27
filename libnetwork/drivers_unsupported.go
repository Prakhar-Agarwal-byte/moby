//go:build !freebsd && !linux && !windows

package libnetwork

import "github.com/Prakhar-Agarwal-byte/moby/libnetwork/driverapi"

func registerNetworkDrivers(r driverapi.Registerer, driverConfig func(string) map[string]interface{}) error {
	return nil
}
