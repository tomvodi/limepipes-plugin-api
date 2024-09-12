package interfaces

import (
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/tune"
	"github.com/tomvodi/limepipes-plugin-api/plugin/v1/messages"
)

type LimePipesPlugin interface {
	PluginInfo() (*messages.PluginInfoResponse, error)

	ParseFromFile(filePath string) ([]*messages.ParsedTune, error)
	Parse(data []byte) ([]*messages.ParsedTune, error)

	ExportToFile(tunes []*tune.Tune, filepath string) error
	Export(tunes []*tune.Tune) ([]byte, error)
}
