package v1

import (
	"context"

	v1 "github.com/brevdev/cloud/v1"
)

const CloudProviderID = "hyperstack"

type HyperstackCredential struct {
	RefID  string
	APIKey string
}

var _ v1.CloudCredential = &HyperstackCredential{}

func NewHyperstackCredential(refID, apiKey string) *HyperstackCredential {
	return &HyperstackCredential{RefID: refID, APIKey: apiKey}
}

func (c *HyperstackCredential) GetReferenceID() string                 { return c.RefID }
func (c *HyperstackCredential) GetAPIType() v1.APIType                 { return v1.APITypeGlobal }
func (c *HyperstackCredential) GetCloudProviderID() v1.CloudProviderID { return CloudProviderID }
func (c *HyperstackCredential) GetTenantID() (string, error)           { return "", nil }

func (c *HyperstackCredential) MakeClient(ctx context.Context, location string) (v1.CloudClient, error) {
	client := NewHyperstackClient(c.RefID, c.APIKey)
	return client.MakeClient(ctx, location)
}

type HyperstackClient struct {
	v1.NotImplCloudClient
	refID    string
	apiKey   string
	location string
}

var _ v1.CloudClient = &HyperstackClient{}

func NewHyperstackClient(refID, apiKey string) *HyperstackClient {
	return &HyperstackClient{refID: refID, apiKey: apiKey}
}

func (c *HyperstackClient) GetAPIType() v1.APIType                 { return v1.APITypeGlobal }
func (c *HyperstackClient) GetCloudProviderID() v1.CloudProviderID { return CloudProviderID }
func (c *HyperstackClient) GetReferenceID() string                 { return c.refID }
func (c *HyperstackClient) GetTenantID() (string, error)           { return "", nil }

func (c *HyperstackClient) MakeClient(_ context.Context, location string) (v1.CloudClient, error) {
	clone := *c
	clone.location = location
	return &clone, nil
}
