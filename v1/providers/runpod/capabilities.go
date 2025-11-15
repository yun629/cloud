package v1

import (
	"context"

	v1 "github.com/brevdev/cloud/v1"
)

func getRunpodCapabilities() v1.Capabilities {
	return v1.Capabilities{}
}

func (c *RunpodClient) GetCapabilities(_ context.Context) (v1.Capabilities, error) {
	return getRunpodCapabilities(), nil
}

func (c *RunpodCredential) GetCapabilities(_ context.Context) (v1.Capabilities, error) {
	return getRunpodCapabilities(), nil
}
