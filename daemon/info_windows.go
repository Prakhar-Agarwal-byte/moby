package daemon // import "github.com/Prakhar-Agarwal-byte/moby/daemon"

import (
	"github.com/Prakhar-Agarwal-byte/moby/api/types"
	"github.com/Prakhar-Agarwal-byte/moby/api/types/system"
	"github.com/Prakhar-Agarwal-byte/moby/daemon/config"
	"github.com/Prakhar-Agarwal-byte/moby/pkg/sysinfo"
)

// fillPlatformInfo fills the platform related info.
func (daemon *Daemon) fillPlatformInfo(v *system.Info, sysInfo *sysinfo.SysInfo, cfg *configStore) {
}

func (daemon *Daemon) fillPlatformVersion(v *types.Version, cfg *configStore) {}

func fillDriverWarnings(v *system.Info) {
}

func cgroupNamespacesEnabled(sysInfo *sysinfo.SysInfo, cfg *config.Config) bool {
	return false
}

// Rootless returns true if daemon is running in rootless mode
func Rootless(*config.Config) bool {
	return false
}

func noNewPrivileges(*config.Config) bool {
	return false
}
