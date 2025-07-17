package infoblox

import (
	"github.com/boltKrank/terraformer/terraformutils"
)

/*
See: https://github.com/infobloxopen/terraform-provider-infoblox/blob/master/infoblox/provider.go


PORT
SSLMODE
CONNECT_TIMEOUT
POOL_CONNECTIONS
WAPI_VERSION
*/

type InfobloxProvider struct {
	terraformutils.Provider
	server           string
	username         string
	password         string
	port             string
	sslmode          string
	connect_timeout  string
	pool_connections string
	wapi_version     string
}

/* AWS example

type AWSProvider struct { //nolint
	terraformutils.Provider
	region  string
	profile string
}

*/

/*
func (p *InfobloxProvider) Init(args []string) error {
	// Handle environment setup or credentials here
	return nil
}

func (p *InfobloxProvider) GetName() string {
	return "infoblox"
}

func (p *InfobloxProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{
		"provider": map[string]interface{}{
			"infoblox": map[string]interface{}{},
		},
	}
}

func (p *InfobloxProvider) GetSupportedService() map[string]terraform.ResourceGenerator {
	return map[string]terraform.ResourceGenerator{
		"record_a": &RecordAGenerator{},
	}
} */
