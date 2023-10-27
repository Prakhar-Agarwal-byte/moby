package libnetwork

import (
	windriver "github.com/Prakhar-Agarwal-byte/moby/libnetwork/drivers/windows"
	"github.com/Prakhar-Agarwal-byte/moby/libnetwork/options"
	"github.com/Prakhar-Agarwal-byte/moby/libnetwork/types"
)

const libnGWNetwork = "nat"

func getPlatformOption() EndpointOption {
	epOption := options.Generic{
		windriver.DisableICC: true,
		windriver.DisableDNS: true,
	}
	return EndpointOptionGeneric(epOption)
}

func (c *Controller) createGWNetwork() (*Network, error) {
	return nil, types.NotImplementedErrorf("default gateway functionality is not implemented in windows")
}
