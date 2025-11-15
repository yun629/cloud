package v1

import (
	"context"
	"fmt"
	"time"

	v1 "github.com/brevdev/cloud/v1"
)

func (c *VastAIClient) GetInstanceTypes(context.Context, v1.GetInstanceTypeArgs) ([]v1.InstanceType, error) {
	return nil, fmt.Errorf("VastAIClient.GetInstanceTypes not implemented")
}

func (c *VastAIClient) GetInstanceTypePollTime() time.Duration {
	return time.Minute
}
