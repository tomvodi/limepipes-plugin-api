package interfaces

import "github.com/tomvodi/limepipes-plugin-api/plugin/v1/messages"

type LimePipesPlugin interface {
	PluginInfo() (*messages.PluginInfoResponse, error)
	ImportFile(filePath string) (*messages.ImportFileResponse, error)
}
