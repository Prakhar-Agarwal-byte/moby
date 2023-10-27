package client // import "github.com/Prakhar-Agarwal-byte/moby/client"

import (
	"context"
	"net/url"

	"github.com/Prakhar-Agarwal-byte/moby/api/types"
)

// PluginDisable disables a plugin
func (cli *Client) PluginDisable(ctx context.Context, name string, options types.PluginDisableOptions) error {
	query := url.Values{}
	if options.Force {
		query.Set("force", "1")
	}
	resp, err := cli.post(ctx, "/plugins/"+name+"/disable", query, nil, nil)
	ensureReaderClosed(resp)
	return err
}
