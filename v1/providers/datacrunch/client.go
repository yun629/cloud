package v1

import (
	"context"

	v1 "github.com/brevdev/cloud/v1"
)

const CloudProviderID = "data-crunch"

type DataCrunchCredential struct {
	RefID  string
	APIKey string
}

var _ v1.CloudCredential = &DataCrunchCredential{}

func NewDataCrunchCredential(refID, apiKey string) *DataCrunchCredential {
	return &DataCrunchCredential{RefID: refID, APIKey: apiKey}
}

func (c *DataCrunchCredential) GetReferenceID() string                 { return c.RefID }
func (c *DataCrunchCredential) GetAPIType() v1.APIType                 { return v1.APITypeGlobal }
func (c *DataCrunchCredential) GetCloudProviderID() v1.CloudProviderID { return CloudProviderID }
func (c *DataCrunchCredential) GetTenantID() (string, error)           { return "", nil }

func (c *DataCrunchCredential) MakeClient(ctx context.Context, location string) (v1.CloudClient, error) {
	client := NewDataCrunchClient(c.RefID, c.APIKey)
	return client.MakeClient(ctx, location)
}

type DataCrunchClient struct {
	v1.NotImplCloudClient
	refID    string
	apiKey   string
	location string
}

var _ v1.CloudClient = &DataCrunchClient{}

func NewDataCrunchClient(refID, apiKey string) *DataCrunchClient {
	return &DataCrunchClient{refID: refID, apiKey: apiKey}
}

func (c *DataCrunchClient) GetAPIType() v1.APIType                 { return v1.APITypeGlobal }
func (c *DataCrunchClient) GetCloudProviderID() v1.CloudProviderID { return CloudProviderID }
func (c *DataCrunchClient) GetReferenceID() string                 { return c.refID }
func (c *DataCrunchClient) GetTenantID() (string, error)           { return "", nil }

func (c *DataCrunchClient) MakeClient(_ context.Context, location string) (v1.CloudClient, error) {
	clone := *c
	clone.location = location
	return &clone, nil
}
