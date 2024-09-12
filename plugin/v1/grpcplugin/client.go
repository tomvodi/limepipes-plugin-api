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

func (c *Client) ExportToFile(
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

func (c *Client) ParseFromFile(
	filePath string,
) ([]*messages.ParsedTune, error) {
	resp, err := c.impl.ParseFromFile(
		context.Background(),
		&messages.ParseFromFileRequest{
			FilePath: filePath,
		},
	)

	if err != nil {
		return nil, err
	}

	return resp.Tunes, nil
}

func (c *Client) Parse(
	data []byte,
) ([]*messages.ParsedTune, error) {
	resp, err := c.impl.Parse(
		context.Background(),
		&messages.ParseRequest{
			Data: data,
		},
	)

	if err != nil {
		return nil, err
	}

	return resp.Tunes, nil
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
