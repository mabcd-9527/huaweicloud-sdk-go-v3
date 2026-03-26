package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HtapCreateLtsConfigRequestBodyLogConfigs struct {

	// **参数解释**：  HTAP标准版实例ID，此参数是实例的唯一标识。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，后缀为in17，长度为36个字符。  **默认取值**：  不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**： 查询日志类型。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**： 不涉及。
	LogType string `json:"log_type"`

	// **参数解释**： LTS日志组ID。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**： 不涉及。
	LtsGroupId string `json:"lts_group_id"`

	// **参数解释**： LTS日志流ID。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**： 不涉及。
	LtsStreamId string `json:"lts_stream_id"`
}

func (o HtapCreateLtsConfigRequestBodyLogConfigs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HtapCreateLtsConfigRequestBodyLogConfigs struct{}"
	}

	return strings.Join([]string{"HtapCreateLtsConfigRequestBodyLogConfigs", string(data)}, " ")
}
