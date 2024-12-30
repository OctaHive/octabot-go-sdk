package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"go.bytecodealliance.org/cm"

	"github.com/OctaHive/octabot-go-sdk/gen/octahive/octabot/plugin"
	wasihttp "github.com/OctaHive/octabot-go-sdk/http"
)

var (
	wasiTransport = &wasihttp.Transport{ConnectTimeout: 30 * time.Second}
	httpClient    = &http.Client{Transport: wasiTransport}
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

		req, err := http.NewRequest(http.MethodGet, "https://dogapi.dog/api/v2/breeds", http.NoBody)
		if err != nil {
			log.Fatal(err)
		}

		res, err := httpClient.Do(req)
		if err != nil {
			panic(err.Error())
		}
		defer res.Body.Close()

		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Status: %d\n", res.StatusCode)
		fmt.Printf("Body: %s\n", string(resBody))

		return cm.OK[cm.Result[plugin.ErrorShape, cm.List[plugin.Action], plugin.Error]](cm.ToList(actions))
	}
}

func main() {}
