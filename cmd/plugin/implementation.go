package main

import (
	"github.com/tomvodi/limepipes-plugin-api/plugin/v1/interfaces"
	"github.com/tomvodi/limepipes-plugin-api/plugin/v1/messages"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type impl struct {
}

func (i impl) ImportLocalFile(filePath string) (*messages.ImportFileResponse, error) {
	return nil, status.Error(codes.Unimplemented, "ImportLocalFile not implemented")
}

func (i impl) Import([]byte) (*messages.ImportFileResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Import not implemented")
}

func (i impl) PluginInfo() (*messages.PluginInfoResponse, error) {
	return &messages.PluginInfoResponse{
		Name:        "Bagpipe Music Writer",
		Description: "A plugin that opens Bagpipe Music Writer and Bagpipe Player files.",
		Type:        messages.PluginType_IN,
		FileTypes:   []string{"bww", "bmw"},
	}, nil
}

func NewImplementation() interfaces.LimePipesPlugin {
	return &impl{}
}
