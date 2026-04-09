package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyEipResponse Response Object
type ModifyEipResponse struct {

	// **参数解释**：  工作流ID。  **取值范围**：  不涉及。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ModifyEipResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyEipResponse struct{}"
	}

	return strings.Join([]string{"ModifyEipResponse", string(data)}, " ")
}
