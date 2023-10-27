package container

import (
	"bytes"
	"context"

	"github.com/Prakhar-Agarwal-byte/moby/api/types"
	"github.com/Prakhar-Agarwal-byte/moby/client"
)

// ExecResult represents a result returned from Exec()
type ExecResult struct {
	ExitCode  int
	outBuffer *bytes.Buffer
	errBuffer *bytes.Buffer
}

// Stdout returns stdout output of a command run by Exec()
func (res *ExecResult) Stdout() string {
	return res.outBuffer.String()
}

// Stderr returns stderr output of a command run by Exec()
func (res *ExecResult) Stderr() string {
	return res.errBuffer.String()
}

// Combined returns combined stdout and stderr output of a command run by Exec()
func (res *ExecResult) Combined() string {
	return res.outBuffer.String() + res.errBuffer.String()
}

// Exec executes a command inside a container, returning the result
// containing stdout, stderr, and exit code. Note:
//   - this is a synchronous operation;
//   - cmd stdin is closed.
func Exec(ctx context.Context, apiClient client.APIClient, id string, cmd []string, ops ...func(*types.ExecConfig)) (ExecResult, error) {
	// prepare exec
	execConfig := types.ExecConfig{
		AttachStdout: true,
		AttachStderr: true,
		Cmd:          cmd,
	}

	for _, op := range ops {
		op(&execConfig)
	}

	cresp, err := apiClient.ContainerExecCreate(ctx, id, execConfig)
	if err != nil {
		return ExecResult{}, err
	}
	execID := cresp.ID

	// run it, with stdout/stderr attached
	aresp, err := apiClient.ContainerExecAttach(ctx, execID, types.ExecStartCheck{})
	if err != nil {
		return ExecResult{}, err
	}

	// read the output
	s, err := demultiplexStreams(ctx, aresp)
	if err != nil {
		return ExecResult{}, err
	}

	// get the exit code
	iresp, err := apiClient.ContainerExecInspect(ctx, execID)
	if err != nil {
		return ExecResult{}, err
	}

	return ExecResult{ExitCode: iresp.ExitCode, outBuffer: &s.stdout, errBuffer: &s.stderr}, nil
}
