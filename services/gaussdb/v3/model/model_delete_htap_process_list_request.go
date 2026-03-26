package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteHtapProcessListRequest Request Object
type DeleteHtapProcessListRequest struct {

	// **参数解释**：  请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us - zh-cn  **默认值**：  en-us。
	XLanguage string `json:"X-Language"`

	// **参数解释**：  HTAP标准版实例ID，严格匹配UUID规则。可通过调用[查询HTAP实例列表](https://support.huaweicloud.com/api-taurusdb/ListHtapInstanceInfo.html)接口获取。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，后缀为in17，且长度为36个字符。  **默认值**：  不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**：  内容类型。  **约束限制**：  不涉及。  **取值范围**：  application/json。  **默认值**：  application/json。
	ContentType string `json:"Content-Type"`

	Body *DeleteHtapProcessReq `json:"body,omitempty"`
}

func (o DeleteHtapProcessListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHtapProcessListRequest struct{}"
	}

	return strings.Join([]string{"DeleteHtapProcessListRequest", string(data)}, " ")
}
