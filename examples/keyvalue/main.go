package main

import (
	"fmt"

	"github.com/OctaHive/octabot-go-sdk/gen/octahive/octabot/plugin"
	"go.bytecodealliance.org/cm"

	"github.com/OctaHive/octabot-go-sdk/gen/wasi/keyvalue/store"
)

var pluginMetadata = plugin.Metadata{
	Name:        "ExamplePlugin",
	Version:     "0.1.0",
	Author:      "Your Name",
	Description: "An example plugin implemented in TinyGo.",
}

func init() {
	plugin.Exports.Load = func() plugin.Metadata {
		return pluginMetadata
	}

	plugin.Exports.Init = func(config string) (result cm.Result[plugin.Error, struct{}, plugin.Error]) {
		return cm.OK[cm.Result[plugin.Error, struct{}, plugin.Error]](struct{}{})
	}

	plugin.Exports.Process = func(payload string) (result cm.Result[plugin.ErrorShape, cm.List[plugin.PluginResult], plugin.Error]) {
		actions := []plugin.PluginResult{}

		bucket, err, isErr := store.Open("").Result()
		if isErr {
			return cm.Err[cm.Result[plugin.ErrorShape, cm.List[plugin.PluginResult], plugin.Error]](plugin.Error(err))
		}

		r := bucket.Set("key", cm.ToList([]uint8("value")))
		if r.IsErr() {
			return cm.Err[cm.Result[plugin.ErrorShape, cm.List[plugin.PluginResult], plugin.Error]](plugin.Error(*r.Err()))
		}

		val, err, isErr := bucket.Get("key").Result()
		if isErr {
			return cm.Err[cm.Result[plugin.ErrorShape, cm.List[plugin.PluginResult], plugin.Error]](plugin.Error(err))
		}

		fmt.Println(string(val.Value().Slice()))

		return cm.OK[cm.Result[plugin.ErrorShape, cm.List[plugin.PluginResult], plugin.Error]](cm.ToList(actions))
	}
}

func main() {}
