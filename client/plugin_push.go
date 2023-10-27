package client // import "github.com/Prakhar-Agarwal-byte/moby/client"

import (
	"context"
	"io"
	"net/http"

	"github.com/Prakhar-Agarwal-byte/moby/api/types/registry"
)

// PluginPush pushes a plugin to a registry
func (cli *Client) PluginPush(ctx context.Context, name string, registryAuth string) (io.ReadCloser, error) {
	resp, err := cli.post(ctx, "/plugins/"+name+"/push", nil, nil, http.Header{
		registry.AuthHeader: {registryAuth},
	})
	if err != nil {
		return nil, err
	}
	return resp.body, nil
}
