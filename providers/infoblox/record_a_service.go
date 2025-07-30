package infoblox

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/boltKrank/terraformer/terraformutils"
)

type ARecord struct {
	Name     string `json:"name"`
	IPv4Addr string `json:"ipv4addr"`
	View     string `json:"view"`
}

type ARecordGenerator struct {
	terraformutils.Service
}

func (g *ARecordGenerator) InitResources() error {
	host := os.Getenv("INFOBLOX_HOST")
	user := os.Getenv("INFOBLOX_USERNAME")
	pass := os.Getenv("INFOBLOX_PASSWORD")

	url := fmt.Sprintf("https://%s/wapi/v2.10/record:a?_return_fields=name,ipv4addr,view", host)

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
		g.Resources = append(g.Resources, terraformutils.NewResource(
			id,
			id,
			"infoblox_a_record",
			"infoblox",
			map[string]string{
				"name":     rec.Name,
				"ipv4addr": rec.IPv4Addr,
				"view":     rec.View,
			},
			[]string{},
			map[string]interface{}{},
		))
	}

	return nil
}
