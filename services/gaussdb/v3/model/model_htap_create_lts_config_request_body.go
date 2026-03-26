package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HtapCreateLtsConfigRequestBody **参数解释**：  打开LTS配置。  **约束限制**：  不涉及。
type HtapCreateLtsConfigRequestBody struct {

	// **参数解释**： 日志配置信息。  **约束限制**： 不涉及。
	LogConfigs []HtapCreateLtsConfigRequestBodyLogConfigs `json:"log_configs"`
}

func (o HtapCreateLtsConfigRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HtapCreateLtsConfigRequestBody struct{}"
	}

	return strings.Join([]string{"HtapCreateLtsConfigRequestBody", string(data)}, " ")
}
