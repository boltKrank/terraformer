package cmd

import (
	"github.com/boltKrank/terraformer/cmd"
	"github.com/boltKrank/terraformer/providers/infoblox"
)

func init() {
	cmd.AddProviderCommand(
		infoblox.NewProvider(),
		"infoblox",
		"Import Infoblox resources using Infoblox WAPI.",
	)
}
