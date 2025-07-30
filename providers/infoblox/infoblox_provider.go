package infoblox

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type InfobloxProvider struct {
	terraformutils.Provider
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
