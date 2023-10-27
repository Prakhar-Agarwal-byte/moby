package client // import "github.com/Prakhar-Agarwal-byte/moby/client"

import (
	"context"
	"net/url"

	"github.com/Prakhar-Agarwal-byte/moby/api/types/swarm"
)

// SecretUpdate attempts to update a secret.
func (cli *Client) SecretUpdate(ctx context.Context, id string, version swarm.Version, secret swarm.SecretSpec) error {
	if err := cli.NewVersionError(ctx, "1.25", "secret update"); err != nil {
		return err
	}
	query := url.Values{}
	query.Set("version", version.String())
	resp, err := cli.post(ctx, "/secrets/"+id+"/update", query, secret, nil)
	ensureReaderClosed(resp)
	return err
}
