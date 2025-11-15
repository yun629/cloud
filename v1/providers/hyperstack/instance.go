package v1

import (
	"context"
	"fmt"

	v1 "github.com/brevdev/cloud/v1"
)

func (c *HyperstackClient) CreateInstance(context.Context, v1.CreateInstanceAttrs) (*v1.Instance, error) {
	return nil, fmt.Errorf("HyperstackClient.CreateInstance not implemented")
}

func (c *HyperstackClient) GetInstance(context.Context, v1.CloudProviderInstanceID) (*v1.Instance, error) {
	return nil, fmt.Errorf("HyperstackClient.GetInstance not implemented")
}

func (c *HyperstackClient) ListInstances(context.Context, v1.ListInstancesArgs) ([]v1.Instance, error) {
	return nil, fmt.Errorf("HyperstackClient.ListInstances not implemented")
}

func (c *HyperstackClient) TerminateInstance(context.Context, v1.CloudProviderInstanceID) error {
	return fmt.Errorf("HyperstackClient.TerminateInstance not implemented")
}

func (c *HyperstackClient) MergeInstanceForUpdate(_ v1.Instance, newInst v1.Instance) v1.Instance {
	return newInst
}

func (c *HyperstackClient) MergeInstanceTypeForUpdate(_ v1.InstanceType, newType v1.InstanceType) v1.InstanceType {
	return newType
}
