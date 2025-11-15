package v1

import (
	"context"

	v1 "github.com/brevdev/cloud/v1"
)

func getHyperstackCapabilities() v1.Capabilities {
	return v1.Capabilities{}
}

func (c *HyperstackClient) GetCapabilities(_ context.Context) (v1.Capabilities, error) {
	return getHyperstackCapabilities(), nil
}

func (c *HyperstackCredential) GetCapabilities(_ context.Context) (v1.Capabilities, error) {
	return getHyperstackCapabilities(), nil
}
