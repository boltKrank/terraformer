package cmd

import (
	"errors"
	"os"

	infoblox_terraformer "github.com/boltKrank/terraformer/providers/infoblox"

	"github.com/boltKrank/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

func newCmdInfobloxImporter(options ImportOptions) *cobra.Command {

	var infoblox_host, username string

	cmd := &cobra.Command{
		Use:   "infoblox",
		Short: "Import current state to Terraform configuration from Infoblox",
		Long:  "Import current state to Terraform configuration from Infoblox",
		RunE: func(cmd *cobra.Command, args []string) error {
			if infoblox_host = os.Getenv("INFOBLOX_HOSTNAME"); infoblox_host == "" {
				return errors.New("requires INFOBLOX_HOSTNAME env var")
			}

			if username == "" {
				return errors.New("requires --user flag to be set")
			}

			provider := newInfobloxProvider()
			err := Import(provider, options, []string{infoblox_host, username})
			if err != nil {
				return err
			}
			return nil
		},
	}

	cmd.AddCommand(listCmd(newInfobloxProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "app,addon", "app=ID")
	cmd.PersistentFlags().StringVarP(&username, "user", "", "", "")
	return cmd
}

func newInfobloxProvider() terraformutils.ProviderGenerator {
	return &infoblox_terraformer.InfobloxProvider{}
}
