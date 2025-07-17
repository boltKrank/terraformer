package infoblox

import (
	"github.com/boltKrank/terraformer/terraformutils"
)

type InfobloxProvider struct {
	terraformutils.Provider
}

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
