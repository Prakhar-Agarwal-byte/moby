//go:build !windows

package homedir // import "github.com/Prakhar-Agarwal-byte/moby/pkg/homedir"

const (
	envKeyName   = "HOME"
	homeShortCut = "~"
)
