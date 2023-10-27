package supervisor // import "github.com/Prakhar-Agarwal-byte/moby/libcontainerd/supervisor"

import "syscall"

// containerdSysProcAttr returns the SysProcAttr to use when exec'ing
// containerd
func containerdSysProcAttr() *syscall.SysProcAttr {
	return nil
}
