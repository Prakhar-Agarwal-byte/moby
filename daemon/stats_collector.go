package daemon // import "github.com/Prakhar-Agarwal-byte/moby/daemon"

import (
	"runtime"
	"time"

	"github.com/Prakhar-Agarwal-byte/moby/daemon/stats"
	"github.com/Prakhar-Agarwal-byte/moby/pkg/meminfo"
)

// newStatsCollector returns a new statsCollector that collections
// stats for a registered container at the specified interval.
// The collector allows non-running containers to be added
// and will start processing stats when they are started.
func (daemon *Daemon) newStatsCollector(interval time.Duration) *stats.Collector {
	// FIXME(vdemeester) move this elsewhere
	if runtime.GOOS == "linux" {
		meminfo, err := meminfo.Read()
		if err == nil && meminfo.MemTotal > 0 {
			daemon.machineMemory = uint64(meminfo.MemTotal)
		}
	}
	s := stats.NewCollector(daemon, interval)
	go s.Run()
	return s
}
