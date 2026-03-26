package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHtapProcessListRequest Request Object
type ShowHtapProcessListRequest struct {

	// **参数解释**：  请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us - zh-cn  **默认值**：  en-us。
	XLanguage string `json:"X-Language"`

	// **参数解释**：  HTAP标准版实例ID，严格匹配UUID规则。可通过调用[查询HTAP实例列表](https://support.huaweicloud.com/api-taurusdb/ListHtapInstanceInfo.html)接口获取。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，后缀为in17，且长度为36个字符。  **默认值**：  不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**：              查询记录数。  **约束限制**：  必须为整数，不能为负数。  **取值范围**：  1-100。  **默认值**：  100。
	Limit *string `json:"limit,omitempty"`

	// **参数解释**：              索引位置，偏移量。从第一条数据偏移offset条数据后开始查询。  **约束限制**：  必须为整数，不能为负数。  **取值范围**：  ≥0  **默认值**：  0。
	Offset *string `json:"offset,omitempty"`
}

func (o ShowHtapProcessListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHtapProcessListRequest struct{}"
	}

	return strings.Join([]string{"ShowHtapProcessListRequest", string(data)}, " ")
}
