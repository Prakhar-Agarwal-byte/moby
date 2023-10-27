//go:build !windows

package daemon // import "github.com/Prakhar-Agarwal-byte/moby/daemon"

import (
	"github.com/Prakhar-Agarwal-byte/moby/container"
	"github.com/Prakhar-Agarwal-byte/moby/pkg/archive"
	"github.com/Prakhar-Agarwal-byte/moby/pkg/idtools"
)

func (daemon *Daemon) tarCopyOptions(container *container.Container, noOverwriteDirNonDir bool) (*archive.TarOptions, error) {
	if container.Config.User == "" {
		return daemon.defaultTarCopyOptions(noOverwriteDirNonDir), nil
	}

	user, err := idtools.LookupUser(container.Config.User)
	if err != nil {
		return nil, err
	}

	identity := idtools.Identity{UID: user.Uid, GID: user.Gid}

	return &archive.TarOptions{
		NoOverwriteDirNonDir: noOverwriteDirNonDir,
		ChownOpts:            &identity,
	}, nil
}
