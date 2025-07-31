package infoblox

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/boltKrank/terraformer/terraformutils"
)

type InfobloxService struct {
	terraformutils.Service
}

func (s *InfobloxService) GenerateClient() string {

	ibHost := s.Args["ibHost"].(string)
	ibUsername := s.Args["ibUsername"].(string)
	ibPassword := s.Args["ibPassword"].(string)
	wapiVersion := s.Args["ibWapiVersion"].(string)

	// Construct the WAPI URL
	baseURL := fmt.Sprintf("https://%s/wapi/%s/", ibHost, wapiVersion)
	endpoint := "record:a?_return_fields=name,ipv4addr"

	fullURL, err := url.JoinPath(baseURL, endpoint)
	if err != nil {
		panic(fmt.Sprintf("Invalid URL: %v", err))
	}

	// Create HTTP client (with timeout)
	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	// Build GET request with Basic Auth
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		panic(fmt.Sprintf("Failed to create request: %v", err))
	}
	req.SetBasicAuth(ibUsername, ibPassword)
	req.Header.Set("Accept", "application/json")

	// Send request
	resp, err := client.Do(req)
	if err != nil {
		panic(fmt.Sprintf("HTTP request failed: %v", err))
	}
	defer resp.Body.Close()

	// Read and print response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(fmt.Sprintf("Failed to read response: %v", err))
	}

	if resp.StatusCode >= 400 {
		panic(fmt.Sprintf("WAPI error: %s", body))
	}

	// Optionally: parse JSON
	var result interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		panic(fmt.Sprintf("Failed to parse JSON: %v", err))
	}

	fmt.Printf("WAPI response:\n%v\n", result)

	return "client"

}
