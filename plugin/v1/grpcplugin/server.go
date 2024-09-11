package grpcplugin

import (
	"context"
	"fmt"
	pluginv1 "github.com/tomvodi/limepipes-plugin-api/plugin/v1"
	"github.com/tomvodi/limepipes-plugin-api/plugin/v1/interfaces"
	"github.com/tomvodi/limepipes-plugin-api/plugin/v1/messages"
)

type Server struct {
	pluginv1.UnimplementedPluginServiceServer
	impl interfaces.LimePipesPlugin
}

//nolint:unused
func (s *Server) mustEmbedUnimplementedPluginServiceServer() {
}

func (s *Server) PluginInfo(
	context.Context,
	*messages.PluginInfoRequest,
) (*messages.PluginInfoResponse, error) {
	return s.impl.PluginInfo()
}

func (s *Server) ImportFile(
	_ context.Context,
	req *messages.ImportFileRequest,
) (*messages.ImportFileResponse, error) {
	switch impType := req.ImportFile.(type) {
	case *messages.ImportFileRequest_FileData:
		return s.impl.Import(impType.FileData)
	case *messages.ImportFileRequest_FilePath:
		return s.impl.ImportLocalFile(impType.FilePath)
	default:
		return nil, fmt.Errorf("unhandled import type %T", impType)
	}
}

func NewGrpcServer(
	impl interfaces.LimePipesPlugin,
) *Server {
	return &Server{
		impl: impl,
	}
}