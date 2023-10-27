//go:build !linux

package archive // import "github.com/Prakhar-Agarwal-byte/moby/pkg/archive"

func getWhiteoutConverter(format WhiteoutFormat, inUserNS bool) (tarWhiteoutConverter, error) {
	return nil, nil
}
