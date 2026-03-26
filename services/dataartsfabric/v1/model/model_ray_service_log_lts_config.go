package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RayServiceLogLtsConfig **参数解释**：RayService日志配置。 **约束限制**：不涉及。
type RayServiceLogLtsConfig struct {

	// **参数解释**：是否开启RayService日志转储到LTS。 **约束限制**：不涉及。 **取值范围**：true，false。 **默认取值**：false。
	Enabled *bool `json:"enabled,omitempty"`

	// **参数解释**：RayService日志转储到的LTS日志组id。 **约束限制**：需要是已经存在的LTS日志组id。 **取值范围**：不涉及。 **默认取值**：不涉及。
	LogGroupId *string `json:"log_group_id,omitempty"`

	// **参数解释**：RayService日志转储到的LTS日志流id。 **约束限制**：需要是已经存在的LTS日志流id。 **取值范围**：不涉及。 **默认取值**：不涉及
	LogStreamId *string `json:"log_stream_id,omitempty"`
}

func (o RayServiceLogLtsConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RayServiceLogLtsConfig struct{}"
	}

	return strings.Join([]string{"RayServiceLogLtsConfig", string(data)}, " ")
}
