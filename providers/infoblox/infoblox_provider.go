package infoblox

import (
	"github.com/boltKrank/terraformer/terraformutils"
)

type InfobloxProvider struct {
	terraformutils.Provider
}

func NewProvider() terraformutils.ProviderGenerator {
	return &InfobloxProvider{}
}

func (p *InfobloxProvider) Init(args []string) error {
	return nil
}

func (p *InfobloxProvider) GetName() string {
	return "infoblox"
}

func (p *InfobloxProvider) GetResourceGenerators() map[string]terraformutils.ResourceGenerator {
	return map[string]terraformutils.ResourceGenerator{
		"record_a": &ARecordGenerator{},
	}
}
