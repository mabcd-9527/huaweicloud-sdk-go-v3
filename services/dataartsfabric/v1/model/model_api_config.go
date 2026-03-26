package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApiConfig - **参数解释**：Ray Service独立Api配置信息。 - **约束限制**：不涉及。 - **取值范围**：不涉及。 - **默认取值**：不涉及。
type ApiConfig struct {

	// - **参数解释**：是否开启Ray Service独立Api功能。 - **约束限制**：不涉及。 - **取值范围**：开启true，关闭false。 - **默认取值**：false。
	EnableIndependentApi *bool `json:"enable_independent_api,omitempty"`

	// - **参数解释**：Ray Service独立Api设置的接口超时时间。 - **约束限制**：不涉及。 - **取值范围**：[0,600000]。 - **默认取值**：60000。
	BackendTimeout *int32 `json:"backend_timeout,omitempty"`
}

func (o ApiConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiConfig struct{}"
	}

	return strings.Join([]string{"ApiConfig", string(data)}, " ")
}
