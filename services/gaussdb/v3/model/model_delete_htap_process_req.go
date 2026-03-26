package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteHtapProcessReq **参数解释**：  删除HTAP实例会话请求体。  **约束限制**：  不涉及。
type DeleteHtapProcessReq struct {

	// **参数解释**：  HTAP实例会话ID列表，可通过调用[查询HTAP实例当前会话](https://support.huaweicloud.com/api-taurusdb/ShowHtapProcessList.html)。  **约束限制**：  实例会话ID小于等于20个。  **取值范围**：  不涉及。  **默认值**：  不涉及。
	ProcessList []string `json:"process_list"`
}

func (o DeleteHtapProcessReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHtapProcessReq struct{}"
	}

	return strings.Join([]string{"DeleteHtapProcessReq", string(data)}, " ")
}
