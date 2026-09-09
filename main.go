package main

import (
	"flag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	fortimanager "github.com/terraform-providers/terraform-provider-fortimanager/fmg"
)

func main() {
	var debugFlag bool

	flag.BoolVar(&debugFlag, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: fortimanager.Provider,
		ProviderAddr: "registry.terraform.io/providers/fortinetdev/fortimanager",
		Debug:        debugFlag,
	})
}
