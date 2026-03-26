package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RayServiceBriefConfig **参数解释**：RayService类型端点的简洁配置。 **约束限制**：不涉及。
type RayServiceBriefConfig struct {
	RayServeConfig *RayServeConfig `json:"ray_serve_config"`

	RayClusterConfig *RayClusterConfig `json:"ray_cluster_config"`

	LogConfig *RayServiceLogConfig `json:"log_config,omitempty"`

	OidcConfig *OidcConfig `json:"oidc_config,omitempty"`

	ApiConfig *ApiConfig `json:"api_config,omitempty"`

	AgencyConfig *AgencyConfig `json:"agency_config,omitempty"`
}

func (o RayServiceBriefConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RayServiceBriefConfig struct{}"
	}

	return strings.Join([]string{"RayServiceBriefConfig", string(data)}, " ")
}
