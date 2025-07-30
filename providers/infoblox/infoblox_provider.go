package infoblox

import (
	"github.com/boltKrank/terraformer/terraformutils"
)

type InfobloxProvider struct {
	terraformutils.Provider
}

// Called by the CLI
func NewProvider() terraformutils.ProviderGenerator {
	return &InfobloxProvider{}
}

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

// Required by interface, but not used in Infoblox
func (p *InfobloxProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{}
}

// ✅ Required to fully satisfy the interface
func (p *InfobloxProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

func (p *InfobloxProvider) InitService(service string, verbose bool) error {
	// You could add logging here if verbose is true
	return nil
}
