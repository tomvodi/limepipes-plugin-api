package main

import (
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"github.com/tomvodi/limepipes-plugin-api/plugin/v1/common"
	"github.com/tomvodi/limepipes-plugin-api/plugin/v1/grpc_plugin"
	"github.com/tomvodi/limepipes-plugin-api/plugin/v1/interfaces"
	"google.golang.org/grpc/status"
	"log"
	"os"
	"os/exec"
)

func main() {
	// Create an hclog.Logger
	logger := hclog.New(&hclog.LoggerOptions{
		Name:   "plugin",
		Output: os.Stdout,
		Level:  hclog.Trace,
	})

	// We're a host. Start by launching the plugin process.
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: common.HandshakeConfig,
		Plugins: plugin.PluginSet{
			"bww": grpc_plugin.NewGrpcPlugin(nil),
		},
		Cmd:              exec.Command("sh", "-c", "./go_build_plugin"),
		AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC},
		Logger:           logger,
	})
	defer client.Kill()

	// Connect via RPC
	rpcClient, err := client.Client()
	if err != nil {
		log.Fatal(err)
	}

	// Request the plugin
	raw, err := rpcClient.Dispense("bww")
	if err != nil {
		log.Fatal(err)
	}

	lpPlugin := raw.(interfaces.LimePipesPlugin)
	info, err := lpPlugin.PluginInfo()
	if err != nil {
		log.Fatal(err)
	}

	logger.Info("Received Plugin info: %v", info)

	logger.Info("Importing file")
	_, err = lpPlugin.ImportFile("test.bww")
	if err != nil {
		state, ok := status.FromError(err)
		if ok {
			logger.Error("Error importing file: %v", state.Message())
			return
		} else {
			logger.Error("Generic error importing file: %v", err)
		}
	}
}
