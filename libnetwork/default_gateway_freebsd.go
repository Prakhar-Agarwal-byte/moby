package libnetwork

import "github.com/Prakhar-Agarwal-byte/moby/libnetwork/types"

const libnGWNetwork = "docker_gwbridge"

func getPlatformOption() EndpointOption {
	return nil
}

func (c *Controller) createGWNetwork() (*Network, error) {
	return nil, types.NotImplementedErrorf("default gateway functionality is not implemented in freebsd")
}
