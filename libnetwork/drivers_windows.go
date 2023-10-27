package libnetwork

import (
	"fmt"

	"github.com/Prakhar-Agarwal-byte/moby/libnetwork/driverapi"
	"github.com/Prakhar-Agarwal-byte/moby/libnetwork/drivers/null"
	"github.com/Prakhar-Agarwal-byte/moby/libnetwork/drivers/windows"
	"github.com/Prakhar-Agarwal-byte/moby/libnetwork/drivers/windows/overlay"
)

func registerNetworkDrivers(r driverapi.Registerer, driverConfig func(string) map[string]interface{}) error {
	for _, nr := range []struct {
		ntype    string
		register func(driverapi.Registerer) error
	}{
		{ntype: null.NetworkType, register: null.Register},
		{ntype: overlay.NetworkType, register: overlay.Register},
	} {
		if err := nr.register(r); err != nil {
			return fmt.Errorf("failed to register %q driver: %w", nr.ntype, err)
		}
	}

	return windows.RegisterBuiltinLocalDrivers(r, driverConfig)
}
