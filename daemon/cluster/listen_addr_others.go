//go:build !linux

package cluster // import "github.com/Prakhar-Agarwal-byte/moby/daemon/cluster"

import "net"

func (c *Cluster) resolveSystemAddr() (net.IP, error) {
	return c.resolveSystemAddrViaSubnetCheck()
}
