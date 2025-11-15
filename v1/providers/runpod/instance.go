package v1

import (
	"context"
	"fmt"

	v1 "github.com/brevdev/cloud/v1"
)

func (c *RunpodClient) CreateInstance(context.Context, v1.CreateInstanceAttrs) (*v1.Instance, error) {
	return nil, fmt.Errorf("RunpodClient.CreateInstance not implemented")
}

func (c *RunpodClient) GetInstance(context.Context, v1.CloudProviderInstanceID) (*v1.Instance, error) {
	return nil, fmt.Errorf("RunpodClient.GetInstance not implemented")
}

func (c *RunpodClient) ListInstances(context.Context, v1.ListInstancesArgs) ([]v1.Instance, error) {
	return nil, fmt.Errorf("RunpodClient.ListInstances not implemented")
}

func (c *RunpodClient) TerminateInstance(context.Context, v1.CloudProviderInstanceID) error {
	return fmt.Errorf("RunpodClient.TerminateInstance not implemented")
}

func (c *RunpodClient) MergeInstanceForUpdate(_ v1.Instance, newInst v1.Instance) v1.Instance {
	return newInst
}

func (c *RunpodClient) MergeInstanceTypeForUpdate(_ v1.InstanceType, newType v1.InstanceType) v1.InstanceType {
	return newType
}
