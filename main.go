package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"terraform-provider-fortimanager/fmg"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: fortimanager.Provider})
}
