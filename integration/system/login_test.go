package system // import "github.com/Prakhar-Agarwal-byte/moby/integration/system"

import (
	"fmt"
	"testing"

	"github.com/Prakhar-Agarwal-byte/moby/api/types/registry"
	"github.com/Prakhar-Agarwal-byte/moby/integration/internal/requirement"
	registrypkg "github.com/Prakhar-Agarwal-byte/moby/registry"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
	"gotest.tools/v3/skip"
)

// Test case for GitHub 22244
func TestLoginFailsWithBadCredentials(t *testing.T) {
	skip.If(t, !requirement.HasHubConnectivity(t))

	ctx := setupTest(t)
	client := testEnv.APIClient()

	config := registry.AuthConfig{
		Username: "no-user",
		Password: "no-password",
	}
	_, err := client.RegistryLogin(ctx, config)
	assert.Assert(t, err != nil)
	assert.Check(t, is.ErrorContains(err, "unauthorized: incorrect username or password"))
	assert.Check(t, is.ErrorContains(err, fmt.Sprintf("https://%s/v2/", registrypkg.DefaultRegistryHost)))
}
