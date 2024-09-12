package grpcplugin

import (
	"context"
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/tune"
	pluginv1 "github.com/tomvodi/limepipes-plugin-api/plugin/v1"
	"github.com/tomvodi/limepipes-plugin-api/plugin/v1/messages"
)

type Client struct {
	impl pluginv1.PluginServiceClient
}

func (c *Client) ExportToLocalFile(
	tunes []*tune.Tune,
	filepath string,
) error {
	_, err := c.impl.ExportToFile(
		context.Background(),
		&messages.ExportToFileRequest{
			ExportTunes: tunes,
			FilePath:    filepath,
		},
	)

	return err
}

func (c *Client) Export(
	tunes []*tune.Tune,
) ([]byte, error) {
	resp, err := c.impl.Export(context.Background(),
		&messages.ExportRequest{
			ExportTunes: tunes,
		},
	)
	if err != nil {
		return nil, err
	}

	return resp.GetData(), err
}

func (c *Client) ImportLocalFile(
	filePath string,
) (*messages.ImportFileResponse, error) {
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
