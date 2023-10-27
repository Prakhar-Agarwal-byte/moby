package client // import "github.com/Prakhar-Agarwal-byte/moby/client"

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/Prakhar-Agarwal-byte/moby/api/types"
	"github.com/Prakhar-Agarwal-byte/moby/api/types/image"
)

// ImageRemove removes an image from the docker host.
func (cli *Client) ImageRemove(ctx context.Context, imageID string, options types.ImageRemoveOptions) ([]image.DeleteResponse, error) {
	query := url.Values{}

	if options.Force {
		query.Set("force", "1")
	}
	if !options.PruneChildren {
		query.Set("noprune", "1")
	}

	var dels []image.DeleteResponse
	resp, err := cli.delete(ctx, "/images/"+imageID, query, nil)
	defer ensureReaderClosed(resp)
	if err != nil {
		return dels, err
	}

	err = json.NewDecoder(resp.body).Decode(&dels)
	return dels, err
}
