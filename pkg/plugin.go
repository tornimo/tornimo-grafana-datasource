package main

import (
	"github.com/hashicorp/go-plugin"
	"github.com/grafana/grafana_plugin_model/go/datasource"
	"net/http"
)

func main() {
	pl := &datasource.DatasourcePluginImpl{
		Plugin: &TornimoDatasource{
			HttpClient: http.DefaultClient,
		},
	}

	plugin.Serve(&plugin.ServeConfig{

		HandshakeConfig: plugin.HandshakeConfig{
			ProtocolVersion:  1,
			MagicCookieKey:   "grafana_plugin_type",
			MagicCookieValue: "datasource",
		},
		Plugins: map[string]plugin.Plugin{
			"backend-datasource": pl,
		},

		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
