//go:build !windows

package main

import (
	"strings"
	"testing"

	"github.com/Prakhar-Agarwal-byte/moby/integration-cli/cli"
	"github.com/Prakhar-Agarwal-byte/moby/testutil"
	"github.com/Prakhar-Agarwal-byte/moby/testutil/request"
	"gotest.tools/v3/assert"
)

// #19100 This is a deprecated feature test, it should be removed in Docker 1.12
func (s *DockerNetworkSuite) TestDeprecatedDockerNetworkStartAPIWithHostconfig(c *testing.T) {
	const netName = "test"
	const conName = "foo"
	cli.DockerCmd(c, "network", "create", netName)
	cli.DockerCmd(c, "create", "--name", conName, "busybox", "top")

	config := map[string]interface{}{
		"HostConfig": map[string]interface{}{
			"NetworkMode": netName,
		},
	}
	_, _, err := request.Post(testutil.GetContext(c), formatV123StartAPIURL("/containers/"+conName+"/start"), request.JSONBody(config))
	assert.NilError(c, err)
	cli.WaitRun(c, conName)
	networks := inspectField(c, conName, "NetworkSettings.Networks")
	assert.Assert(c, strings.Contains(networks, netName), "Should contain '%s' network", netName)
	assert.Assert(c, !strings.Contains(networks, "bridge"), "Should not contain 'bridge' network")
}
