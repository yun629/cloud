package v1

import (
	"context"

	v1 "github.com/brevdev/cloud/v1"
)

func getVastAICapabilities() v1.Capabilities {
	return v1.Capabilities{}
}

func (c *VastAIClient) GetCapabilities(_ context.Context) (v1.Capabilities, error) {
	return getVastAICapabilities(), nil
}

func (c *VastAICredential) GetCapabilities(_ context.Context) (v1.Capabilities, error) {
	return getVastAICapabilities(), nil
}
