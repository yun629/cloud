package v1

import (
	"context"

	v1 "github.com/brevdev/cloud/v1"
)

func getDataCrunchCapabilities() v1.Capabilities {
	return v1.Capabilities{}
}

func (c *DataCrunchClient) GetCapabilities(_ context.Context) (v1.Capabilities, error) {
	return getDataCrunchCapabilities(), nil
}

func (c *DataCrunchCredential) GetCapabilities(_ context.Context) (v1.Capabilities, error) {
	return getDataCrunchCapabilities(), nil
}
