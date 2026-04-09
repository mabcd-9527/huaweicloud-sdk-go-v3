package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RollbackMigrationResponse Response Object
type RollbackMigrationResponse struct {

	// **参数解释**：  工作流ID。  **取值范围**：  不涉及。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RollbackMigrationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RollbackMigrationResponse struct{}"
	}

	return strings.Join([]string{"RollbackMigrationResponse", string(data)}, " ")
}
