package fileutils // import "github.com/Prakhar-Agarwal-byte/moby/pkg/fileutils"

// GetTotalUsedFds Returns the number of used File Descriptors. Not supported
// on Windows.
func GetTotalUsedFds() int {
	return -1
}
