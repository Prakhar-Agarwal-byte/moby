package tarexport // import "github.com/Prakhar-Agarwal-byte/moby/image/tarexport"

import (
	"github.com/Prakhar-Agarwal-byte/moby/api/types/events"
	"github.com/Prakhar-Agarwal-byte/moby/image"
	"github.com/Prakhar-Agarwal-byte/moby/layer"
	refstore "github.com/Prakhar-Agarwal-byte/moby/reference"
	"github.com/docker/distribution"
)

const (
	manifestFileName           = "manifest.json"
	legacyLayerFileName        = "layer.tar"
	legacyConfigFileName       = "json"
	legacyRepositoriesFileName = "repositories"
)

type manifestItem struct {
	Config       string
	RepoTags     []string
	Layers       []string
	Parent       image.ID                                 `json:",omitempty"`
	LayerSources map[layer.DiffID]distribution.Descriptor `json:",omitempty"`
}

type tarexporter struct {
	is             image.Store
	lss            layer.Store
	rs             refstore.Store
	loggerImgEvent LogImageEvent
}

// LogImageEvent defines interface for event generation related to image tar(load and save) operations
type LogImageEvent interface {
	// LogImageEvent generates an event related to an image operation
	LogImageEvent(imageID, refName string, action events.Action)
}

// NewTarExporter returns new Exporter for tar packages
func NewTarExporter(is image.Store, lss layer.Store, rs refstore.Store, loggerImgEvent LogImageEvent) image.Exporter {
	return &tarexporter{
		is:             is,
		lss:            lss,
		rs:             rs,
		loggerImgEvent: loggerImgEvent,
	}
}
