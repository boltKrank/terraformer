package infoblox

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/boltKrank/terraformer/terraformutils"
)

type ARecord struct {
	Name         string            `json:"name"`
	IPv4Addr     string            `json:"ipv4addr"`
	View         string            `json:"view"`
	TTL          *int              `json:"ttl,omitempty"`
	Comment      string            `json:"comment,omitempty"`
	ExtAttrs     map[string]string `json:"extattrs,omitempty"`
	ExtAttrsJSON string            // compute JSON string if needed
}

type ARecordGenerator struct {
	terraformutils.Service
}

// InitResources is required by terraformutils.ServiceGenerator
func (g *ARecordGenerator) InitResources() error {
	host := os.Getenv("INFOBLOX_HOST")
	user := os.Getenv("INFOBLOX_USERNAME")
	pass := os.Getenv("INFOBLOX_PASSWORD")

	url := fmt.Sprintf("https://%s/wapi/v2.13/record:a", host)

	req, _ := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(user, pass)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var result struct {
		Records []ARecord `json:"result"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}

	for _, rec := range result.Records {
		id := fmt.Sprintf("%s_%s", rec.Name, rec.IPv4Addr)
		attrs := map[string]string{
			"fqdn":    rec.Name,
			"address": rec.IPv4Addr,
			"view":    rec.View,
		}
		if rec.TTL != nil {
			attrs["ttl"] = fmt.Sprintf("%d", *rec.TTL)
		}
		if rec.Comment != "" {
			attrs["comment"] = rec.Comment
		}
		// Optionally include ext_attrs if available
		if len(rec.ExtAttrs) > 0 {
			attrs["ext_attrs"] = rec.ExtAttrsJSON
		}

		g.Resources = append(g.Resources, terraformutils.NewResource(
			id,
			id,
			"infoblox_a_record",
			"infoblox",
			attrs,
			[]string{},
			map[string]interface{}{},
		))
	}

	return nil
}
