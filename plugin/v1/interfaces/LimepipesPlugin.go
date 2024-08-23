package interfaces

import "github.com/tomvodi/limepipes-plugin-api/plugin/v1/messages"

type LimePipesPlugin interface {
	PluginInfo() (*messages.PluginInfoResponse, error)
	ImportLocalFile(filePath string) (*messages.ImportFileResponse, error)
	Import(data []byte) (*messages.ImportFileResponse, error)
}
