package grpcplugin

import (
	"context"
	"github.com/hashicorp/go-plugin"
	pluginv1 "github.com/tomvodi/limepipes-plugin-api/plugin/v1"
	"github.com/tomvodi/limepipes-plugin-api/plugin/v1/interfaces"
	"google.golang.org/grpc"
)

type GrpcPlugin struct {
	plugin.Plugin
	impl interfaces.LimePipesPlugin
}

func (g *GrpcPlugin) GRPCServer(
	_ *plugin.GRPCBroker,
	server *grpc.Server,
) error {
	var s pluginv1.PluginServiceServer = NewGrpcServer(g.impl)
	pluginv1.RegisterPluginServiceServer(server, s)
	return nil
}

func (g *GrpcPlugin) GRPCClient(
	_ context.Context,
	_ *plugin.GRPCBroker,
	conn *grpc.ClientConn,
) (any, error) {
	client := pluginv1.NewPluginServiceClient(conn)

	var c interfaces.LimePipesPlugin = NewGrpcClient(client)

	return c, nil
}

func NewGrpcPlugin(
	impl interfaces.LimePipesPlugin,
) *GrpcPlugin {
	return &GrpcPlugin{
		impl: impl,
	}
}
