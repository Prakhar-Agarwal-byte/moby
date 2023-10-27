package system // import "github.com/Prakhar-Agarwal-byte/moby/pkg/system"

import "errors"

// ErrNotSupportedPlatform means the platform is not supported.
var ErrNotSupportedPlatform = errors.New("platform and architecture is not supported")
