package libnetwork

import (
	"github.com/Prakhar-Agarwal-byte/moby/libnetwork/driverapi"
	"github.com/Prakhar-Agarwal-byte/moby/libnetwork/drivers/null"
)

func registerNetworkDrivers(r driverapi.Registerer, driverConfig func(string) map[string]interface{}) error {
	return null.Register(r)
}
