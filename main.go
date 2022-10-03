package main

import (
	"context"
	"flag"

	tsm "terraform-provider-tsm/plugin/provider"

	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	var debugMode bool

	flag.BoolVar(&debugMode, "debuggable", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	if debugMode {
		err := plugin.Debug(context.Background(), "provider",
			&plugin.ServeOpts{
				ProviderFunc: tsm.Provider,
			})
		if err != nil {
			//log.Println(err.Error())
		}
	} else {
		plugin.Serve(&plugin.ServeOpts{
			ProviderFunc: tsm.Provider})
	}

}
