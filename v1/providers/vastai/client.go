package v1

import (
	"context"

	v1 "github.com/brevdev/cloud/v1"
)

const CloudProviderID = "vast-ai"

// VastAICredential implements CloudCredential for Vast.ai.
type VastAICredential struct {
	RefID  string
	APIKey string
}

var _ v1.CloudCredential = &VastAICredential{}

func NewVastAICredential(refID, apiKey string) *VastAICredential {
	return &VastAICredential{RefID: refID, APIKey: apiKey}
}

func (c *VastAICredential) GetReferenceID() string                 { return c.RefID }
func (c *VastAICredential) GetAPIType() v1.APIType                 { return v1.APITypeGlobal }
func (c *VastAICredential) GetCloudProviderID() v1.CloudProviderID { return CloudProviderID }
func (c *VastAICredential) GetTenantID() (string, error)           { return "", nil }

func (c *VastAICredential) MakeClient(ctx context.Context, location string) (v1.CloudClient, error) {
	client := NewVastAIClient(c.RefID, c.APIKey)
	return client.MakeClient(ctx, location)
}

// VastAIClient implements CloudClient for Vast.ai.
type VastAIClient struct {
	v1.NotImplCloudClient
	refID    string
	apiKey   string
	location string
}

var _ v1.CloudClient = &VastAIClient{}

func NewVastAIClient(refID, apiKey string) *VastAIClient {
	return &VastAIClient{refID: refID, apiKey: apiKey}
}

func (c *VastAIClient) GetAPIType() v1.APIType                 { return v1.APITypeGlobal }
func (c *VastAIClient) GetCloudProviderID() v1.CloudProviderID { return CloudProviderID }
func (c *VastAIClient) GetReferenceID() string                 { return c.refID }
func (c *VastAIClient) GetTenantID() (string, error)           { return "", nil }

func (c *VastAIClient) MakeClient(_ context.Context, location string) (v1.CloudClient, error) {
	clone := *c
	clone.location = location
	return &clone, nil
}
