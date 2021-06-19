package hibp

import (
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/schema"
)

type HibpConfig struct {
	ApiKey *string `cty:"apiKey"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"apiKey": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &HibpConfig{}
}

func GetConfig(connection *plugin.Connection) HibpConfig {
	if connection == nil || connection.Config == nil {
		return HibpConfig{}
	}

	config, _ := connection.Config.(HibpConfig)
	return config
}
