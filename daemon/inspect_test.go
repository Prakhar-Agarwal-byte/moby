package daemon // import "github.com/Prakhar-Agarwal-byte/moby/daemon"

import (
	"testing"

	containertypes "github.com/Prakhar-Agarwal-byte/moby/api/types/container"
	"github.com/Prakhar-Agarwal-byte/moby/container"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestGetInspectData(t *testing.T) {
	c := &container.Container{
		ID:           "inspect-me",
		HostConfig:   &containertypes.HostConfig{},
		State:        container.NewState(),
		ExecCommands: container.NewExecStore(),
	}

	d := &Daemon{
		linkIndex: newLinkIndex(),
	}
	if d.UsesSnapshotter() {
		t.Skip("does not apply to containerd snapshotters, which don't have RWLayer set")
	}
	cfg := &configStore{}
	d.configStore.Store(cfg)

	_, err := d.getInspectData(&cfg.Config, c)
	assert.Check(t, is.ErrorContains(err, "RWLayer of container inspect-me is unexpectedly nil"))

	c.Dead = true
	_, err = d.getInspectData(&cfg.Config, c)
	assert.Check(t, err)
}
