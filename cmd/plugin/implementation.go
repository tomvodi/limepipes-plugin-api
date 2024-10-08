package main

import (
	"github.com/tomvodi/limepipes-plugin-api/musicmodel/v1/tune"
	"github.com/tomvodi/limepipes-plugin-api/plugin/v1/fileformat"
	"github.com/tomvodi/limepipes-plugin-api/plugin/v1/messages"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Impl struct {
}

func (i *Impl) ExportToFile(_ []*tune.Tune, _ string) error {
	return status.Error(codes.Unimplemented, "ExportToLocalFile not implemented")
}

func (i *Impl) Export(
	_ []*tune.Tune,
) ([]byte, error) {
	return nil, status.Error(codes.Unimplemented, "Export not implemented")
}

func (i *Impl) ParseFromFile(_ string) ([]*messages.ParsedTune, error) {
	return nil, status.Error(codes.Unimplemented, "ImportLocalFile not implemented")
}

func (i *Impl) Parse(_ []byte) ([]*messages.ParsedTune, error) {
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
