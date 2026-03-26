package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHtapProcessListResponse Response Object
type ShowHtapProcessListResponse struct {

	// **参数解释**：  HTAP标准版查询的会话信息。  **默认值**：  不涉及。
	ProcessList    *[]HtapProcessInfo `json:"process_list,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ShowHtapProcessListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHtapProcessListResponse struct{}"
	}

	return strings.Join([]string{"ShowHtapProcessListResponse", string(data)}, " ")
}
