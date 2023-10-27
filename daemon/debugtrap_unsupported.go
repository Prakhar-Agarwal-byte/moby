//go:build !linux && !darwin && !freebsd && !windows

package daemon // import "github.com/Prakhar-Agarwal-byte/moby/daemon"

func (daemon *Daemon) setupDumpStackTrap(_ string) {
	return
}
