package daemon // import "github.com/Prakhar-Agarwal-byte/moby/daemon"

import "github.com/Prakhar-Agarwal-byte/moby/daemon/config"

// reloadPlatform updates configuration with platform specific options
// and updates the passed attributes
func (daemon *Daemon) reloadPlatform(txn *reloadTxn, newCfg *configStore, conf *config.Config, attributes map[string]string) error {
	return nil
}
