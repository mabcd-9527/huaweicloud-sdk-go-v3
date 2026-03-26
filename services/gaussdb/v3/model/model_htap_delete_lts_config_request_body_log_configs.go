package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HtapDeleteLtsConfigRequestBodyLogConfigs struct {

	// **参数解释**：  HTAP标准版实例ID，此参数是实例的唯一标识。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，后缀为in17，长度为36个字符。  **默认取值**：  不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**： 日志类型。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**： 不涉及。
	LogType string `json:"log_type"`
}

func (o HtapDeleteLtsConfigRequestBodyLogConfigs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HtapDeleteLtsConfigRequestBodyLogConfigs struct{}"
	}

	return strings.Join([]string{"HtapDeleteLtsConfigRequestBodyLogConfigs", string(data)}, " ")
}
