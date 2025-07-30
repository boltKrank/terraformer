package cmd

import (
	"github.com/GoogleCloudPlatform/terraformer/cmd"
	"github.com/boltKrank/terraformer/providers/infoblox"
)

func init() {
	cmd.AddProviderCommand(
		infoblox.NewProvider(), // your custom provider
		"infoblox",             // provider name
		"Import Infoblox resources using Infoblox WAPI.",
	)
}
