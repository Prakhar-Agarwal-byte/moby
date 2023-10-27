//go:build unix && !linux

package chrootarchive // import "github.com/Prakhar-Agarwal-byte/moby/pkg/chrootarchive"

import (
	"testing"

	"github.com/Prakhar-Agarwal-byte/moby/pkg/reexec"
)

func TestMain(m *testing.M) {
	if reexec.Init() {
		return
	}
	m.Run()
}
