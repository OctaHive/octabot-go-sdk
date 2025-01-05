package main

import (
	"go.bytecodealliance.org/cm"

	"github.com/OctaHive/octabot-go-sdk/gen/octahive/octabot/plugin"
	"github.com/OctaHive/octabot-go-sdk/log"
)

// Plugin metadata
var pluginMetadata = plugin.Metadata{
	Name:        "ExamplePlugin",
	Version:     "0.1.0",
	Author:      "Your Name",
	Description: "An example plugin implemented in TinyGo.",
}

func init() {
	logger := log.ContextLogger("hello-world")

	plugin.Exports.Load = func() plugin.Metadata {
		return pluginMetadata
	}

	plugin.Exports.Init = func(config string) (result cm.Result[plugin.Error, struct{}, plugin.Error]) {
		logger.Info("Hello World init callback")
		return cm.OK[cm.Result[plugin.Error, struct{}, plugin.Error]](struct{}{})
	}

	plugin.Exports.Process = func(payload string) (result cm.Result[plugin.ErrorShape, cm.List[plugin.PluginResult], plugin.Error]) {
		actions := []plugin.PluginResult{}

		logger.Info("Hello World process callback")

		return cm.OK[cm.Result[plugin.ErrorShape, cm.List[plugin.PluginResult], plugin.Error]](cm.ToList(actions))
	}
}

func main() {}
