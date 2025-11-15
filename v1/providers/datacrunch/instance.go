package v1

import (
	"context"
	"fmt"

	v1 "github.com/brevdev/cloud/v1"
)

func (c *DataCrunchClient) CreateInstance(context.Context, v1.CreateInstanceAttrs) (*v1.Instance, error) {
	return nil, fmt.Errorf("DataCrunchClient.CreateInstance not implemented")
}

func (c *DataCrunchClient) GetInstance(context.Context, v1.CloudProviderInstanceID) (*v1.Instance, error) {
	return nil, fmt.Errorf("DataCrunchClient.GetInstance not implemented")
}

func (c *DataCrunchClient) ListInstances(context.Context, v1.ListInstancesArgs) ([]v1.Instance, error) {
	return nil, fmt.Errorf("DataCrunchClient.ListInstances not implemented")
}

func (c *DataCrunchClient) TerminateInstance(context.Context, v1.CloudProviderInstanceID) error {
	return fmt.Errorf("DataCrunchClient.TerminateInstance not implemented")
}

func (c *DataCrunchClient) MergeInstanceForUpdate(_ v1.Instance, newInst v1.Instance) v1.Instance {
	return newInst
}

func (c *DataCrunchClient) MergeInstanceTypeForUpdate(_ v1.InstanceType, newType v1.InstanceType) v1.InstanceType {
	return newType
}
