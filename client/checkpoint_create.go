package client // import "github.com/Prakhar-Agarwal-byte/moby/client"

import (
	"context"

	"github.com/Prakhar-Agarwal-byte/moby/api/types/checkpoint"
)

// CheckpointCreate creates a checkpoint from the given container with the given name
func (cli *Client) CheckpointCreate(ctx context.Context, container string, options checkpoint.CreateOptions) error {
	resp, err := cli.post(ctx, "/containers/"+container+"/checkpoints", nil, options, nil)
	ensureReaderClosed(resp)
	return err
}
