package grpcplugin

import (
	"context"
	pluginv1 "github.com/tomvodi/limepipes-plugin-api/plugin/v1"
	"github.com/tomvodi/limepipes-plugin-api/plugin/v1/messages"
)

type Client struct {
	impl pluginv1.PluginServiceClient
}

func (c *Client) ImportLocalFile(filePath string) (*messages.ImportFileResponse, error) {
	return c.impl.ImportFile(
		context.Background(),
		&messages.ImportFileRequest{
			ImportFile: &messages.ImportFileRequest_FilePath{
				FilePath: filePath,
			},
		},
	)
}

func (c *Client) Import(fileData []byte) (*messages.ImportFileResponse, error) {
	return c.impl.ImportFile(
		context.Background(),
		&messages.ImportFileRequest{
			ImportFile: &messages.ImportFileRequest_FileData{
				FileData: fileData,
			},
		},
	)
}

func (c *Client) PluginInfo() (
	*messages.PluginInfoResponse,
	error,
) {
	return c.impl.PluginInfo(
		context.Background(),
		&messages.PluginInfoRequest{},
	)
}

func NewGrpcClient(
	impl pluginv1.PluginServiceClient,
) *Client {
	return &Client{
		impl: impl,
	}
}
