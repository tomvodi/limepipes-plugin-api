package grpc_plugin

import (
	"context"
	pluginv1 "github.com/tomvodi/limepipes-plugin-api/plugin/v1"
	"github.com/tomvodi/limepipes-plugin-api/plugin/v1/interfaces"
	"github.com/tomvodi/limepipes-plugin-api/plugin/v1/messages"
)

type grpcClient struct {
	impl pluginv1.PluginServiceClient
}

func (g *grpcClient) ImportFile(filePath string) (*messages.ImportFileResponse, error) {
	return g.impl.ImportFile(
		context.Background(),
		&messages.ImportFileRequest{FilePath: filePath},
	)
}

func (g *grpcClient) PluginInfo() (
	*messages.PluginInfoResponse,
	error,
) {
	return g.impl.PluginInfo(
		context.Background(),
		&messages.PluginInfoRequest{},
	)
}

func NewGrpcClient(
	impl pluginv1.PluginServiceClient,
) interfaces.LimePipesPlugin {
	return &grpcClient{
		impl: impl,
	}
}
