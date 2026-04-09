package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetReadOnlyStatusResponse Response Object
type SetReadOnlyStatusResponse struct {

	// **参数解释**：  工作流ID。  **取值范围**：  不涉及。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetReadOnlyStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetReadOnlyStatusResponse struct{}"
	}

	return strings.Join([]string{"SetReadOnlyStatusResponse", string(data)}, " ")
}
