package main

import (
	"steampipe-plugin-hibp/hibp"

	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: hibp.Plugin})
}
