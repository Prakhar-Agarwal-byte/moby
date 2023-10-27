//go:build !wasip1

package system // import "github.com/Prakhar-Agarwal-byte/moby/pkg/system"

import (
	"io/fs"

	"golang.org/x/sys/unix"
)

// Lgetxattr retrieves the value of the extended attribute identified by attr
// and associated with the given path in the file system.
// It will returns a nil slice and nil error if the xattr is not set.
func Lgetxattr(path string, attr string) ([]byte, error) {
	pathErr := func(err error) ([]byte, error) {
		return nil, &fs.PathError{Op: "lgetxattr", Path: path, Err: err}
	}

	// Start with a 128 length byte array
	dest := make([]byte, 128)
	sz, errno := unix.Lgetxattr(path, attr, dest)

	for errno == unix.ERANGE {
		// Buffer too small, use zero-sized buffer to get the actual size
		sz, errno = unix.Lgetxattr(path, attr, []byte{})
		if errno != nil {
			return pathErr(errno)
		}
		dest = make([]byte, sz)
		sz, errno = unix.Lgetxattr(path, attr, dest)
	}

	switch {
	case errno == unix.ENODATA:
		return nil, nil
	case errno != nil:
		return pathErr(errno)
	}

	return dest[:sz], nil
}

// Lsetxattr sets the value of the extended attribute identified by attr
// and associated with the given path in the file system.
func Lsetxattr(path string, attr string, data []byte, flags int) error {
	err := unix.Lsetxattr(path, attr, data, flags)
	if err != nil {
		return &fs.PathError{Op: "lsetxattr", Path: path, Err: err}
	}
	return nil
}
