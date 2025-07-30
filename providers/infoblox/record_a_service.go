package infoblox

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/boltKrank/terraformer/terraformutils"
)

type ARecord struct {
	Name     string                 `json:"name"`
	IPv4Addr string                 `json:"ipv4addr"`
	View     string                 `json:"view"`
	TTL      *int                   `json:"ttl,omitempty"`
	Comment  string                 `json:"comment,omitempty"`
	ExtAttrs map[string]interface{} `json:"extattrs,omitempty"`
}

type ARecordGenerator struct {
	terraformutils.Service
}

func (g *ARecordGenerator) InitResources() error {
	host := os.Getenv("INFOBLOX_HOST")
	user := os.Getenv("INFOBLOX_USERNAME")
	pass := os.Getenv("INFOBLOX_PASSWORD")

	if host == "" || user == "" || pass == "" {
		return fmt.Errorf("Missing required environment variables: INFOBLOX_HOST, INFOBLOX_USERNAME, INFOBLOX_PASSWORD")
	}

	url := fmt.Sprintf("https://%s/wapi/v2.10/record:a?_return_fields=name,ipv4addr,view,ttl,comment,extattrs", host)

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
			attrs["ttl"] = strconv.Itoa(*rec.TTL)
		}
		if rec.Comment != "" {
			attrs["comment"] = rec.Comment
		}

		for key, val := range rec.ExtAttrs {
			if v, ok := val.(map[string]interface{}); ok {
				if rawVal, ok := v["value"].(string); ok {
					attrs[fmt.Sprintf("extattr_%s", key)] = rawVal
				}
			}
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
