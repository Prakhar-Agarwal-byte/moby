package checkpoint // import "github.com/Prakhar-Agarwal-byte/moby/api/server/router/checkpoint"

import "github.com/Prakhar-Agarwal-byte/moby/api/types/checkpoint"

// Backend for Checkpoint
type Backend interface {
	CheckpointCreate(container string, config checkpoint.CreateOptions) error
	CheckpointDelete(container string, config checkpoint.DeleteOptions) error
	CheckpointList(container string, config checkpoint.ListOptions) ([]checkpoint.Summary, error)
}
