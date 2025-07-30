package infoblox

import (
	"github.com/boltKrank/terraformer/terraformutils"
)

type InfobloxProvider struct {
	terraformutils.Provider
}

// Required to satisfy terraformutils.ProviderGenerator
func (p *InfobloxProvider) Init(args []string) error {
	return nil
}

func (p *InfobloxProvider) GetName() string {
	return "infoblox"
}

func (p *InfobloxProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"record_a": &ARecordGenerator{},
	}
}

// ✅ Required in boltKrank/terraformer for interface compliance
func (p *InfobloxProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{}
}
