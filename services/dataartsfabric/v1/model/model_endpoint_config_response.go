package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EndpointConfigResponse **参数解释**：端点的配置。 **约束限制**：不涉及。
type EndpointConfigResponse struct {
	RayServiceConfig *RayServiceConfigResponse `json:"ray_service_config,omitempty"`
}

func (o EndpointConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EndpointConfigResponse struct{}"
	}

	return strings.Join([]string{"EndpointConfigResponse", string(data)}, " ")
}
