//go:build windows

package daemon // import "github.com/Prakhar-Agarwal-byte/moby/daemon"

import (
	"github.com/Prakhar-Agarwal-byte/moby/daemon/config"
	"github.com/Prakhar-Agarwal-byte/moby/pkg/plugingetter"
)

func registerMetricsPluginCallback(getter plugingetter.PluginGetter, sockPath string) {
}

func (daemon *Daemon) listenMetricsSock(*config.Config) (string, error) {
	return "", nil
}
