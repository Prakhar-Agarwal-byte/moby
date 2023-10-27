package client // import "github.com/Prakhar-Agarwal-byte/moby/client"

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/Prakhar-Agarwal-byte/moby/api/types/system"
)

// Info returns information about the docker server.
func (cli *Client) Info(ctx context.Context) (system.Info, error) {
	var info system.Info
	serverResp, err := cli.get(ctx, "/info", url.Values{}, nil)
	defer ensureReaderClosed(serverResp)
	if err != nil {
		return info, err
	}

	if err := json.NewDecoder(serverResp.body).Decode(&info); err != nil {
		return info, fmt.Errorf("Error reading remote info: %v", err)
	}

	return info, nil
}
