package v1

import (
	"context"
	"fmt"
	"time"

	v1 "github.com/brevdev/cloud/v1"
)

func (c *HyperstackClient) GetInstanceTypes(context.Context, v1.GetInstanceTypeArgs) ([]v1.InstanceType, error) {
	return nil, fmt.Errorf("HyperstackClient.GetInstanceTypes not implemented")
}

func (c *HyperstackClient) GetInstanceTypePollTime() time.Duration {
	return time.Minute
}
