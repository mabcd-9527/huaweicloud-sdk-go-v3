package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EndpointConfig **参数解释**：端点的配置。 **约束限制**：不涉及。
type EndpointConfig struct {
	RayServiceConfig *RayServiceConfig `json:"ray_service_config,omitempty"`
}

func (o EndpointConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EndpointConfig struct{}"
	}

	return strings.Join([]string{"EndpointConfig", string(data)}, " ")
}
