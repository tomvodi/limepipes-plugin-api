package main

import (
	"github.com/tomvodi/limepipes-plugin-api/plugin/v1/fileformat"
	"github.com/tomvodi/limepipes-plugin-api/plugin/v1/messages"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Impl struct {
}

func (i *Impl) ImportLocalFile(string) (*messages.ImportFileResponse, error) {
	return nil, status.Error(codes.Unimplemented, "ImportLocalFile not implemented")
}

func (i *Impl) Import([]byte) (*messages.ImportFileResponse, error) {
	return nil, status.Error(codes.Unimplemented, "Import not implemented")
}

func (i *Impl) PluginInfo() (*messages.PluginInfoResponse, error) {
	return &messages.PluginInfoResponse{
		Name:        "Bagpipe Music Writer",
		Description: "A plugin that opens Bagpipe Music Writer and Bagpipe Player files.",
		Type:        messages.PluginType_IN,
		FileFormat:  fileformat.Format_BWW,
	}, nil
}

func NewImplementation() *Impl {
	return &Impl{}
}
