package grpc_plugin

import (
	"context"
	"fmt"
	pluginv1 "github.com/tomvodi/limepipes-plugin-api/plugin/v1"
	"github.com/tomvodi/limepipes-plugin-api/plugin/v1/interfaces"
	"github.com/tomvodi/limepipes-plugin-api/plugin/v1/messages"
)

type grpcServer struct {
	pluginv1.UnimplementedPluginServiceServer
	impl interfaces.LimePipesPlugin
}

func (g *grpcServer) mustEmbedUnimplementedPluginServiceServer() {
}

func (g *grpcServer) PluginInfo(
	context.Context,
	*messages.PluginInfoRequest,
) (*messages.PluginInfoResponse, error) {
	return g.impl.PluginInfo()
}

func (g *grpcServer) ImportFile(
	_ context.Context,
	req *messages.ImportFileRequest,
) (*messages.ImportFileResponse, error) {
	switch impType := req.ImportFile.(type) {
	case *messages.ImportFileRequest_FileData:
		return g.impl.Import(impType.FileData)
	case *messages.ImportFileRequest_FilePath:
		return g.impl.ImportLocalFile(impType.FilePath)
	default:
		return nil, fmt.Errorf("unhandled import type %T", impType)
	}
}

func NewGrpcServer(
	impl interfaces.LimePipesPlugin,
) pluginv1.PluginServiceServer {
	return &grpcServer{
		impl: impl,
	}
}
