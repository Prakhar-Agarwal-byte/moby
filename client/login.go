package client // import "github.com/Prakhar-Agarwal-byte/moby/client"

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/Prakhar-Agarwal-byte/moby/api/types/registry"
)

// RegistryLogin authenticates the docker server with a given docker registry.
// It returns unauthorizedError when the authentication fails.
func (cli *Client) RegistryLogin(ctx context.Context, auth registry.AuthConfig) (registry.AuthenticateOKBody, error) {
	resp, err := cli.post(ctx, "/auth", url.Values{}, auth, nil)
	defer ensureReaderClosed(resp)

	if err != nil {
		return registry.AuthenticateOKBody{}, err
	}

	var response registry.AuthenticateOKBody
	err = json.NewDecoder(resp.body).Decode(&response)
	return response, err
}
