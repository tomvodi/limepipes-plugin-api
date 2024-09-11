package main

import (
	"github.com/hashicorp/go-plugin"
	"github.com/tomvodi/limepipes-plugin-api/internal/utils"
	"github.com/tomvodi/limepipes-plugin-api/plugin/v1/common"
	"github.com/tomvodi/limepipes-plugin-api/plugin/v1/fileformat"
	"github.com/tomvodi/limepipes-plugin-api/plugin/v1/grpcplugin"
	"google.golang.org/grpc"
)

// defaultGRPCServer returns a new gRPC server with the given options.
// Acts as a factory method for gRPC servers.
func defaultGRPCServer(opts []grpc.ServerOption) *grpc.Server {
	return grpc.NewServer(opts...)
}

func main() {
	utils.SetupConsoleLogger()

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: common.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			fileformat.Format_BWW.String(): grpcplugin.NewGrpcPlugin(NewImplementation()),
		},

		GRPCServer: defaultGRPCServer,
	})
}
