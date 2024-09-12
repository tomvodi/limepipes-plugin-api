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
func (s *Server) mustEmbedUnimplementedPluginServiceServer() {}

func (s *Server) PluginInfo(
	_ context.Context,
	_ *messages.PluginInfoRequest,
) (*messages.PluginInfoResponse, error) {
	return s.impl.PluginInfo()
}

func (s *Server) ParseFromFile(
	_ context.Context,
	req *messages.ParseFromFileRequest,
) (*messages.ParseFromFileResponse, error) {
	if req.FilePath == "" {
		return nil, fmt.Errorf("parse from file needs a non empty file path")
	}

	tunes, err := s.impl.ParseFromFile(req.FilePath)
	if err != nil {
		return nil, err
	}

	return &messages.ParseFromFileResponse{
		Tunes: tunes,
	}, nil
}

func (s *Server) Parse(
	_ context.Context,
	req *messages.ParseRequest,
) (*messages.ParseResponse, error) {
	if len(req.Data) == 0 {
		return nil, fmt.Errorf("parse needs non empty data")
	}

	tunes, err := s.impl.Parse(req.Data)
	if err != nil {
		return nil, err
	}

	return &messages.ParseResponse{
		Tunes: tunes,
	}, nil
}

func (s *Server) Export(
	_ context.Context,
	req *messages.ExportRequest,
) (*messages.ExportResponse, error) {
	data, err := s.impl.Export(req.ExportTunes)
	return &messages.ExportResponse{
		Data: data,
	}, err
}

func (s *Server) ExportToFile(
	_ context.Context,
	req *messages.ExportToFileRequest,
) (*messages.ExportToFileResponse, error) {
	if req.FilePath != "" {
		return &messages.ExportToFileResponse{},
			fmt.Errorf("export to file needs a non empty file path")
	}

	return &messages.ExportToFileResponse{},
		s.impl.ExportToFile(req.ExportTunes, req.FilePath)
}

func NewGrpcServer(
	impl interfaces.LimePipesPlugin,
) *Server {
	return &Server{
		impl: impl,
	}
}
