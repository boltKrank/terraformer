package cmd

import (
	infoblox_terraforming "github.com/boltKrank/terraformer/providers/infoblox"
	"github.com/boltKrank/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

func newCmdInfobloxImporter(options ImportOptions) *cobra.Command {

	var infobloxHost, username, password string

	cmd := &cobra.Command{
		Use:   "infoblox",
		Short: "Import current state to Terraform configuration from Infoblox",
		Long:  "Import current state to Terraform configuration from Infoblox",
		RunE: func(cmd *cobra.Command, args []string) error {
			provider := newInfobloxProvider()
			err := Import(provider, options, []string{infobloxHost, username, password})
			if err != nil {
				return err
			}
			return nil

		},
	}

	cmd.AddCommand(listCmd(newInfobloxProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "infoblox", "")
	cmd.PersistentFlags().StringVarP(&infobloxHost, "infoblox-host", "", "", "INFOBLOX_HOST or env param INFOBLOX_HOST")
	cmd.PersistentFlags().StringVarP(&username, "username", "", "", "INFOBLOX_USERNAME or env param INFOBLOX_USERNAME")
	cmd.PersistentFlags().StringVarP(&password, "password", "", "", "INFOBLOX_PASSWORD or env param INFOBLOX_PASSWORD")
	return cmd
}

func newInfobloxProvider() terraformutils.ProviderGenerator {
	return &infoblox_terraforming.InfobloxProvider{}
}
