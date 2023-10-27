package client // import "github.com/Prakhar-Agarwal-byte/moby/client"

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/Prakhar-Agarwal-byte/moby/errdefs"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestPluginSetError(t *testing.T) {
	client := &Client{
		client: newMockClient(errorMock(http.StatusInternalServerError, "Server error")),
	}

	err := client.PluginSet(context.Background(), "plugin_name", []string{})
	assert.Check(t, is.ErrorType(err, errdefs.IsSystem))
}

func TestPluginSet(t *testing.T) {
	expectedURL := "/plugins/plugin_name/set"

	client := &Client{
		client: newMockClient(func(req *http.Request) (*http.Response, error) {
			if !strings.HasPrefix(req.URL.Path, expectedURL) {
				return nil, fmt.Errorf("Expected URL '%s', got '%s'", expectedURL, req.URL)
			}
			if req.Method != http.MethodPost {
				return nil, fmt.Errorf("expected POST method, got %s", req.Method)
			}
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader([]byte(""))),
			}, nil
		}),
	}

	err := client.PluginSet(context.Background(), "plugin_name", []string{"arg1"})
	if err != nil {
		t.Fatal(err)
	}
}
