package environment // import "github.com/Prakhar-Agarwal-byte/moby/integration-cli/environment"

import (
	"context"
	"os"
	"os/exec"

	"github.com/Prakhar-Agarwal-byte/moby/testutil/environment"
)

// DefaultClientBinary is the name of the docker binary
var DefaultClientBinary = os.Getenv("TEST_CLIENT_BINARY")

func init() {
	if DefaultClientBinary == "" {
		DefaultClientBinary = "docker"
	}
}

// Execution contains information about the current test execution and daemon
// under test
type Execution struct {
	environment.Execution
	dockerBinary string
}

// DockerBinary returns the docker binary for this testing environment
func (e *Execution) DockerBinary() string {
	return e.dockerBinary
}

// New returns details about the testing environment
func New(ctx context.Context) (*Execution, error) {
	env, err := environment.New(ctx)
	if err != nil {
		return nil, err
	}

	dockerBinary, err := exec.LookPath(DefaultClientBinary)
	if err != nil {
		return nil, err
	}

	return &Execution{
		Execution:    *env,
		dockerBinary: dockerBinary,
	}, nil
}
