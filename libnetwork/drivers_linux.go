package libnetwork

import (
	"fmt"

	"github.com/Prakhar-Agarwal-byte/moby/libnetwork/driverapi"
	"github.com/Prakhar-Agarwal-byte/moby/libnetwork/drivers/bridge"
	"github.com/Prakhar-Agarwal-byte/moby/libnetwork/drivers/host"
	"github.com/Prakhar-Agarwal-byte/moby/libnetwork/drivers/ipvlan"
	"github.com/Prakhar-Agarwal-byte/moby/libnetwork/drivers/macvlan"
	"github.com/Prakhar-Agarwal-byte/moby/libnetwork/drivers/null"
	"github.com/Prakhar-Agarwal-byte/moby/libnetwork/drivers/overlay"
)

func registerNetworkDrivers(r driverapi.Registerer, driverConfig func(string) map[string]interface{}) error {
	noConfig := func(fn func(driverapi.Registerer) error) func(driverapi.Registerer, map[string]interface{}) error {
		return func(r driverapi.Registerer, _ map[string]interface{}) error { return fn(r) }
	}

	for _, nr := range []struct {
		ntype    string
		register func(driverapi.Registerer, map[string]interface{}) error
	}{
		{ntype: bridge.NetworkType, register: bridge.Register},
		{ntype: host.NetworkType, register: noConfig(host.Register)},
		{ntype: ipvlan.NetworkType, register: ipvlan.Register},
		{ntype: macvlan.NetworkType, register: macvlan.Register},
		{ntype: null.NetworkType, register: noConfig(null.Register)},
		{ntype: overlay.NetworkType, register: overlay.Register},
	} {
		if err := nr.register(r, driverConfig(nr.ntype)); err != nil {
			return fmt.Errorf("failed to register %q driver: %w", nr.ntype, err)
		}
	}

	return nil
}
