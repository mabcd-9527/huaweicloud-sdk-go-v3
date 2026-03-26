package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HtapDeleteLtsConfigRequestBody **参数解释**：  删除LTS配置。  **约束限制**：  不涉及。
type HtapDeleteLtsConfigRequestBody struct {

	// **参数解释**： LTS配置。  **约束限制**： 不涉及。
	LogConfigs []HtapDeleteLtsConfigRequestBodyLogConfigs `json:"log_configs"`
}

func (o HtapDeleteLtsConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HtapDeleteLtsConfigRequestBody struct{}"
	}

	return strings.Join([]string{"HtapDeleteLtsConfigRequestBody", string(data)}, " ")
}
