package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelMigrationResponse Response Object
type CancelMigrationResponse struct {

	// **参数解释**：  工作流ID。  **取值范围**：  不涉及。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CancelMigrationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelMigrationResponse struct{}"
	}

	return strings.Join([]string{"CancelMigrationResponse", string(data)}, " ")
}
