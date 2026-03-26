package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlarmSubDetailResopnse 告警详细信息
type AlarmSubDetailResopnse struct {

	// **参数解释**：  告警配置ID。  **取值范围**： 不涉及。
	Id string `json:"id"`

	// **参数解释**：  告警定义ID。  **取值范围**： 不涉及。
	AlarmId string `json:"alarm_id"`

	// **参数解释**：  告警名称。  **取值范围**： 不涉及。
	AlarmName string `json:"alarm_name"`

	// **参数解释**：  所属服务。  **取值范围**： 不涉及。
	NameSpace string `json:"name_space"`

	// **参数解释**：  告警级别。  **取值范围**：  不涉及。
	AlarmLevel string `json:"alarm_level"`
}

func (o AlarmSubDetailResopnse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmSubDetailResopnse struct{}"
	}

	return strings.Join([]string{"AlarmSubDetailResopnse", string(data)}, " ")
}
