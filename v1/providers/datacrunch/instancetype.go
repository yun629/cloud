package v1

import (
	"context"
	"fmt"
	"time"

	v1 "github.com/brevdev/cloud/v1"
)

func (c *DataCrunchClient) GetInstanceTypes(context.Context, v1.GetInstanceTypeArgs) ([]v1.InstanceType, error) {
	return nil, fmt.Errorf("DataCrunchClient.GetInstanceTypes not implemented")
}

func (c *DataCrunchClient) GetInstanceTypePollTime() time.Duration {
	return time.Minute
}
