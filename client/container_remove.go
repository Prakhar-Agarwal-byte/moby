package client // import "github.com/Prakhar-Agarwal-byte/moby/client"

import (
	"context"
	"net/url"

	"github.com/Prakhar-Agarwal-byte/moby/api/types/container"
)

// ContainerRemove kills and removes a container from the docker host.
func (cli *Client) ContainerRemove(ctx context.Context, containerID string, options container.RemoveOptions) error {
	query := url.Values{}
	if options.RemoveVolumes {
		query.Set("v", "1")
	}
	if options.RemoveLinks {
		query.Set("link", "1")
	}

	if options.Force {
		query.Set("force", "1")
	}

	resp, err := cli.delete(ctx, "/containers/"+containerID, query, nil)
	defer ensureReaderClosed(resp)
	return err
}
