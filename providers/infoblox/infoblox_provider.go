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
		"infoblox_a_record": &ARecordGenerator{},

		/* TODO:
		ResourcesMap: map[string]*schema.Resource{
			"infoblox_network_view":           resourceNetworkView(),
			"infoblox_ipv4_network_container": resourceIPv4NetworkContainer(),
			"infoblox_ipv6_network_container": resourceIPv6NetworkContainer(),
			"infoblox_ipv4_network":           resourceIPv4Network(),
			"infoblox_ipv6_network":           resourceIPv6Network(),
			"infoblox_ip_allocation":          resourceIPAllocation(),
			"infoblox_ip_association":         resourceIpAssociationInit(),
			"infoblox_a_record":               resourceARecord(),
			"infoblox_aaaa_record":            resourceAAAARecord(),
			"infoblox_cname_record":           resourceCNAMERecord(),
			"infoblox_ptr_record":             resourcePTRRecord(),
			"infoblox_zone_delegated":         resourceZoneDelegated(),
			"infoblox_txt_record":             resourceTXTRecord(),
			"infoblox_mx_record":              resourceMXRecord(),
			"infoblox_srv_record":             resourceSRVRecord(),
			"infoblox_dns_view":               resourceDNSView(),
			"infoblox_zone_auth":              resourceZoneAuth(),
			"infoblox_zone_forward":           resourceZoneForward(),
			"infoblox_dtc_lbdn":               resourceDtcLbdnRecord(),
			"infoblox_dtc_pool":               resourceDtcPool(),
			"infoblox_dtc_server":             resourceDtcServer(),
			"infoblox_ipv4_fixed_address":     resourceFixedRecord(),
			"infoblox_alias_record":           resourceAliasRecord(),
			"infoblox_ns_record":              resourceNSRecord(),
			"infoblox_ipv4_range":             resourceRange(),
			"infoblox_ipv4_range_template":    resourceRangeTemplate(),
			"infoblox_ipv4_shared_network":    resourceIpv4SharedNetwork(),

			*/
		},
	}
}

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
