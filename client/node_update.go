package client // import "github.com/Prakhar-Agarwal-byte/moby/client"

import (
	"context"
	"net/url"

	"github.com/Prakhar-Agarwal-byte/moby/api/types/swarm"
)

// NodeUpdate updates a Node.
func (cli *Client) NodeUpdate(ctx context.Context, nodeID string, version swarm.Version, node swarm.NodeSpec) error {
	query := url.Values{}
	query.Set("version", version.String())
	resp, err := cli.post(ctx, "/nodes/"+nodeID+"/update", query, node, nil)
	ensureReaderClosed(resp)
	return err
}
