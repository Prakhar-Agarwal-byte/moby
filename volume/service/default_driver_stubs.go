//go:build !linux && !windows

package service // import "github.com/Prakhar-Agarwal-byte/moby/volume/service"

import (
	"github.com/Prakhar-Agarwal-byte/moby/pkg/idtools"
	"github.com/Prakhar-Agarwal-byte/moby/volume/drivers"
)

func setupDefaultDriver(_ *drivers.Store, _ string, _ idtools.Identity) error { return nil }
