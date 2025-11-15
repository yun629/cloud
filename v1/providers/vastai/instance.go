package v1

import (
	"context"
	"fmt"

	v1 "github.com/brevdev/cloud/v1"
)

func (c *VastAIClient) CreateInstance(context.Context, v1.CreateInstanceAttrs) (*v1.Instance, error) {
	return nil, fmt.Errorf("VastAIClient.CreateInstance not implemented")
}

func (c *VastAIClient) GetInstance(context.Context, v1.CloudProviderInstanceID) (*v1.Instance, error) {
	return nil, fmt.Errorf("VastAIClient.GetInstance not implemented")
}

func (c *VastAIClient) ListInstances(context.Context, v1.ListInstancesArgs) ([]v1.Instance, error) {
	return nil, fmt.Errorf("VastAIClient.ListInstances not implemented")
}

func (c *VastAIClient) TerminateInstance(context.Context, v1.CloudProviderInstanceID) error {
	return fmt.Errorf("VastAIClient.TerminateInstance not implemented")
}

func (c *VastAIClient) MergeInstanceForUpdate(_ v1.Instance, newInst v1.Instance) v1.Instance {
	return newInst
}

func (c *VastAIClient) MergeInstanceTypeForUpdate(_ v1.InstanceType, newType v1.InstanceType) v1.InstanceType {
	return newType
}
