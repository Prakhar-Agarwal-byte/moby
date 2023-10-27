package containerfs // import "github.com/Prakhar-Agarwal-byte/moby/pkg/containerfs"

import "os"

// EnsureRemoveAll is an alias to os.RemoveAll on Windows
var EnsureRemoveAll = os.RemoveAll
