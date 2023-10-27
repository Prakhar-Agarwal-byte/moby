package libnetwork

import (
	"github.com/Prakhar-Agarwal-byte/moby/libnetwork/ipamapi"
	builtinIpam "github.com/Prakhar-Agarwal-byte/moby/libnetwork/ipams/builtin"
	nullIpam "github.com/Prakhar-Agarwal-byte/moby/libnetwork/ipams/null"
	remoteIpam "github.com/Prakhar-Agarwal-byte/moby/libnetwork/ipams/remote"
	"github.com/Prakhar-Agarwal-byte/moby/libnetwork/ipamutils"
	"github.com/Prakhar-Agarwal-byte/moby/pkg/plugingetter"
)

func initIPAMDrivers(r ipamapi.Registerer, pg plugingetter.PluginGetter, addressPool []*ipamutils.NetworkToSplit) error {
	// TODO: pass address pools as arguments to builtinIpam.Init instead of
	// indirectly through global mutable state. Swarmkit references that
	// function so changing its signature breaks the build.
	if err := builtinIpam.SetDefaultIPAddressPool(addressPool); err != nil {
		return err
	}

	for _, fn := range [](func(ipamapi.Registerer) error){
		builtinIpam.Register,
		nullIpam.Register,
	} {
		if err := fn(r); err != nil {
			return err
		}
	}

	return remoteIpam.Register(r, pg)
}
