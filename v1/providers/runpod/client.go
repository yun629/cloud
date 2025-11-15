package v1

import (
	"context"

	v1 "github.com/brevdev/cloud/v1"
)

const CloudProviderID = "runpod"

type RunpodCredential struct {
	RefID  string
	APIKey string
}

var _ v1.CloudCredential = &RunpodCredential{}

func NewRunpodCredential(refID, apiKey string) *RunpodCredential {
	return &RunpodCredential{RefID: refID, APIKey: apiKey}
}

func (c *RunpodCredential) GetReferenceID() string                 { return c.RefID }
func (c *RunpodCredential) GetAPIType() v1.APIType                 { return v1.APITypeGlobal }
func (c *RunpodCredential) GetCloudProviderID() v1.CloudProviderID { return CloudProviderID }
func (c *RunpodCredential) GetTenantID() (string, error)           { return "", nil }

func (c *RunpodCredential) MakeClient(ctx context.Context, location string) (v1.CloudClient, error) {
	client := NewRunpodClient(c.RefID, c.APIKey)
	return client.MakeClient(ctx, location)
}

type RunpodClient struct {
	v1.NotImplCloudClient
	refID    string
	apiKey   string
	location string
}

var _ v1.CloudClient = &RunpodClient{}

func NewRunpodClient(refID, apiKey string) *RunpodClient {
	return &RunpodClient{refID: refID, apiKey: apiKey}
}

func (c *RunpodClient) GetAPIType() v1.APIType                 { return v1.APITypeGlobal }
func (c *RunpodClient) GetCloudProviderID() v1.CloudProviderID { return CloudProviderID }
func (c *RunpodClient) GetReferenceID() string                 { return c.refID }
func (c *RunpodClient) GetTenantID() (string, error)           { return "", nil }

func (c *RunpodClient) MakeClient(_ context.Context, location string) (v1.CloudClient, error) {
	clone := *c
	clone.location = location
	return &clone, nil
}
