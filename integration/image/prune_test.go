package image

import (
	"testing"

	"github.com/Prakhar-Agarwal-byte/moby/api/types/filters"
	"github.com/Prakhar-Agarwal-byte/moby/integration/internal/container"
	"github.com/Prakhar-Agarwal-byte/moby/testutil/environment"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
	"gotest.tools/v3/skip"
)

// Regression test for: https://github.com/moby/moby/issues/45732
func TestPruneDontDeleteUsedDangling(t *testing.T) {
	skip.If(t, testEnv.DaemonInfo.OSType == "windows", "FIXME: hack/make/.build-empty-images doesn't run on Windows")

	ctx := setupTest(t)
	client := testEnv.APIClient()

	danglingID := environment.GetTestDanglingImageId(testEnv)

	container.Create(ctx, t, client,
		container.WithImage(danglingID),
		container.WithCmd("sleep", "60"))

	pruned, err := client.ImagesPrune(ctx, filters.NewArgs(filters.Arg("dangling", "true")))

	assert.NilError(t, err)
	assert.Check(t, is.Len(pruned.ImagesDeleted, 0))
}
