package main

import (
	"fmt"

	"go.bytecodealliance.org/cm"

	"github.com/OctaHive/octabot-go-sdk/gen/octahive/octabot/plugin"
)

// Plugin metadata
var pluginMetadata = plugin.Metadata{
	Name:        "ExamplePlugin",
	Version:     "0.1.0",
	Author:      "Your Name",
	Description: "An example plugin implemented in TinyGo.",
}

func init() {
	plugin.Exports.Init = func() plugin.Metadata {
		return pluginMetadata
	}

	plugin.Exports.Process = func(config string, payload string) (result cm.Result[plugin.ErrorShape, cm.List[plugin.Action], plugin.Error]) {
		actions := []plugin.Action{}

		fmt.Println("Hello, World!")

		return cm.OK[cm.Result[plugin.ErrorShape, cm.List[plugin.Action], plugin.Error]](cm.ToList(actions))
	}
}

func main() {}
