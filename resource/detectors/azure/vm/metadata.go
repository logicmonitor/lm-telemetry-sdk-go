package vm

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	// Azure IMDS compute endpoint, see https://aka.ms/azureimds
	metadataEndpoint = "http://169.254.169.254/metadata/instance/compute"
)

// Provider gets metadata from the Azure IMDS
type Provider interface {
	Metadata(context.Context) (*ComputeMetadata, error)
}

// ComputeMetadata is the Azure IMDS compute metadata response format
type ComputeMetadata struct {
	Location          string `json:"location"`
	Name              string `json:"name"`
	VMID              string `json:"vmID"`
	VMSize            string `json:"vmSize"`
	SubscriptionID    string `json:"subscriptionID"`
	ResourceGroupName string `json:"resourceGroupName"`
	VMScaleSetName    string `json:"vmScaleSetName"`
}

type azureProviderImpl struct {
	endpoint string
	client   *http.Client
}

// NewProvider creates a new metadata provider
func NewProvider() Provider {
	return &azureProviderImpl{
		endpoint: metadataEndpoint,
		client:   &http.Client{},
	}
}

// Metadata queries a given endpoint and parses the output to the Azure IMDS format
func (p *azureProviderImpl) Metadata(ctx context.Context) (*ComputeMetadata, error) {
	const (
		// API version used
		apiVersionKey = "api-version"
		apiVersion    = "2020-09-01"

		// format used
		formatKey  = "format"
		jsonFormat = "json"
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, p.endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Add("Metadata", "True")
	q := req.URL.Query()
	q.Add(formatKey, jsonFormat)
	q.Add(apiVersionKey, apiVersion)
	req.URL.RawQuery = q.Encode()

	resp, err := p.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to query Azure IMDS: %v", err)
	} else if resp.StatusCode == http.StatusNotFound {
		return nil, errors.New("not executing Azure VM")
	} else if resp.StatusCode != 200 {
		//lint:ignore ST1005 Azure is a capitalized proper noun here
		return nil, fmt.Errorf("Azure IMDS replied with status code: %s", resp.Status)
	}

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read Azure IMDS reply: %v", err)
	}

	var metadata *ComputeMetadata
	err = json.Unmarshal(respBody, &metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to decode Azure IMDS reply: %v", err)
	}

	return metadata, nil
}
