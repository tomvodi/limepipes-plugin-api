package grpc_plugin

import (
	"context"
	"github.com/hashicorp/go-plugin"
	pluginv1 "github.com/tomvodi/limepipes-plugin-api/plugin/v1"
	"github.com/tomvodi/limepipes-plugin-api/plugin/v1/interfaces"
	"google.golang.org/grpc"
)

type grpcPlugin struct {
	plugin.Plugin
	impl interfaces.LimePipesPlugin
}

func (g *grpcPlugin) GRPCServer(
	_ *plugin.GRPCBroker,
	server *grpc.Server,
) error {
	pluginv1.RegisterPluginServiceServer(server, NewGrpcServer(g.impl))
	return nil
}

func (g *grpcPlugin) GRPCClient(
	_ context.Context,
	_ *plugin.GRPCBroker,
	conn *grpc.ClientConn,
) (interface{}, error) {
	client := pluginv1.NewPluginServiceClient(conn)
	return NewGrpcClient(client), nil
}

func NewGrpcPlugin(
	impl interfaces.LimePipesPlugin,
) plugin.Plugin {
	return &grpcPlugin{
		impl: impl,
	}
}
