package client // import "github.com/Prakhar-Agarwal-byte/moby/client"

import (
	"context"

	"github.com/Prakhar-Agarwal-byte/moby/api/types/swarm"
)

// SwarmJoin joins the swarm.
func (cli *Client) SwarmJoin(ctx context.Context, req swarm.JoinRequest) error {
	resp, err := cli.post(ctx, "/swarm/join", nil, req, nil)
	ensureReaderClosed(resp)
	return err
}
