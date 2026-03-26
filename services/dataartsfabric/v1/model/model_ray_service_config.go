package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RayServiceConfig **参数解释**：RayService类型端点的配置。 **约束限制**：不涉及。
type RayServiceConfig struct {
	RayServeConfig *RayServeConfig `json:"ray_serve_config"`

	RayClusterConfig *RayClusterConfig `json:"ray_cluster_config"`

	LogConfig *RayServiceLogConfig `json:"log_config,omitempty"`

	OidcConfig *OidcConfig `json:"oidc_config,omitempty"`

	ApiConfig *ApiConfig `json:"api_config,omitempty"`

	AgencyConfig *AgencyConfig `json:"agency_config,omitempty"`

	// **参数解释**：数据信息。 **约束限制**：不涉及。 **取值范围**：[0,1]。 **默认取值**：不涉及。
	DataInfos *[]DataBriefInfo `json:"data_infos,omitempty"`
}

func (o RayServiceConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RayServiceConfig struct{}"
	}

	return strings.Join([]string{"RayServiceConfig", string(data)}, " ")
}
