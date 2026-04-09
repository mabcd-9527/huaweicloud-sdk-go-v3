package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteGroupResponse Response Object
type DeleteGroupResponse struct {

	// **参数解释**：  工作流ID。  **取值范围**：  不涉及。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGroupResponse struct{}"
	}

	return strings.Join([]string{"DeleteGroupResponse", string(data)}, " ")
}
