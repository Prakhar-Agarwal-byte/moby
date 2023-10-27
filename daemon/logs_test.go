package daemon // import "github.com/Prakhar-Agarwal-byte/moby/daemon"

import (
	"testing"

	containertypes "github.com/Prakhar-Agarwal-byte/moby/api/types/container"
)

func TestMergeAndVerifyLogConfigNilConfig(t *testing.T) {
	d := &Daemon{defaultLogConfig: containertypes.LogConfig{Type: "json-file", Config: map[string]string{"max-file": "1"}}}
	cfg := containertypes.LogConfig{Type: d.defaultLogConfig.Type}
	if err := d.mergeAndVerifyLogConfig(&cfg); err != nil {
		t.Fatal(err)
	}
}
