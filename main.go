package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/terraform-providers/terraform-provider-fortimanager/fmg"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: fortimanager.Provider})
}
