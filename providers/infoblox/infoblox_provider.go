package infoblox

import (
	"errors"
	"os"

	"github.com/boltKrank/terraformer/terraformutils"
	"github.com/zclconf/go-cty/cty"
)

type InfobloxProvider struct {
	terraformutils.Provider
	ibHost        string
	ibUsername    string
	ibPassword    string
	ibWapiVersion string
}

/* Called by the CLI
func NewProvider() terraformutils.ProviderGenerator {
	return &InfobloxProvider{}
} */

func (p *InfobloxProvider) Init(args []string) error {
	ibHost := os.Getenv("INFOBLOX_HOST")
	if ibHost == "" {
		return errors.New("set INFOBLOX_HOST env var")
	}
	p.ibHost = ibHost

	ibUsername := os.Getenv("INFOBLOX_USERNAME")
	if ibUsername == "" {
		return errors.New("set INFOBLOX_USERNAME env var")
	}
	p.ibUsername = ibUsername

	ibPassword := os.Getenv("INFOBLOX_PASSWORD")
	if ibPassword == "" {
		return errors.New("set INFOBLOX_PASSWORD env var")
	}
	p.ibPassword = ibPassword

	return nil
}

func (p *InfobloxProvider) GetName() string {
	return "infoblox"
}

// Required by interface, but not used in Infoblox
func (p *InfobloxProvider) GetProviderData(arg ...string) map[string]interface{} {
	return map[string]interface{}{}
}

// ✅ Required to fully satisfy the interface
func (p *InfobloxProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

func (p *InfobloxProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"infoblox_a_record": 		&ARecordGenerator{},
		"infoblox_cname_record":    &CnameRecordGenerator{},
		"infoblox_aaaa_record":     &AAAARecordGenerator{},
		"infoblox_network_view":    &NetworkViewGenerator{},

		/* TODO:
		ResourcesMap: map[string]*schema.Resource{
			
			"infoblox_ipv4_network_container": &Ipv4NetworkContainerGenerator{},
			"infoblox_ipv6_network_container": &Ipv6NetworkContainerGenerator{},
			"infoblox_ipv4_network":           &Ipv4NetworkGenerator{},
			"infoblox_ipv6_network":           &Ipv6NetworkGenerator{},
			"infoblox_ip_allocation":          &IpAllocationGenerator{},
			"infoblox_ip_association":         &IpAssociationGenerator{},
			
			
			"infoblox_ptr_record":             &PtrRecordGenerator{},
			"infoblox_zone_delegated":         &ZoneDelegatedGenerator{},
			"infoblox_txt_record":             &TxtRecordGenerator{},
			"infoblox_mx_record":              &MxRecordGenerator{},
			"infoblox_srv_record":             &SrvRecordGenerator{},
			"infoblox_dns_view":               &DnsViewGenerator{},
			"infoblox_zone_auth":              &ZoneAuthGenerator{},
			"infoblox_zone_forward":           &ZoneForwardGenerator{},
			"infoblox_dtc_lbdn":               &DtcLbdnGenerator{},
			"infoblox_dtc_pool":               &DtcPoolGenerator{},
			"infoblox_dtc_server":             &DtcServerGenerator{},
			"infoblox_ipv4_fixed_address":     &Ipv4FixedAddressGenerator{},
			"infoblox_alias_record":           &AliasRecordGenerator{},
			"infoblox_ns_record":              &NsRecordGenerator{},
			"infoblox_ipv4_range":             &Ipv4RangeGenerator{},
			"infoblox_ipv4_range_template":    &Ipv4RangeTemplateGenerator{},
			"infoblox_ipv4_shared_network":    &Ipv4SharedNetworkGenerator{},

			*/
		},
	}
}z

func (p *InfobloxProvider) InitService(serviceName string, verbose bool) error {
	// You could add logging here if verbose is true
	var isSupported bool
	if _, isSupported = p.GetSupportedService()[serviceName]; !isSupported {
		return errors.New("infoblox: " + serviceName + " not supported service")
	}
	p.Service = p.GetSupportedService()[serviceName]
	p.Service.SetName(serviceName)
	p.Service.SetVerbose(verbose)
	p.Service.SetProviderName(p.GetName())
	p.Service.SetArgs(map[string]interface{}{
		"ibHost":        p.ibHost,
		"ibUsername":    p.ibUsername,
		"ibPassword":    p.ibPassword,
		"ibWapiVersion": p.ibWapiVersion,
	})
z
	return nil
}

func (p *InfobloxProvider) GetConfig() cty.Value {
	return cty.ObjectVal(map[string]cty.Value{
		"ibHost":        cty.StringVal(p.ibHost),
		"ibUsername":    cty.StringVal(p.ibUsername),
		"ibPassword":    cty.StringVal(p.ibPassword),
		"ibWapiVersion": cty.StringVal(p.ibWapiVersion),
	})
}
