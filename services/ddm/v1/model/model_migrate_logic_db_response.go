package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MigrateLogicDbResponse Response Object
type MigrateLogicDbResponse struct {

	// **参数解释**：  工作流ID。  **取值范围**：  不涉及。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o MigrateLogicDbResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateLogicDbResponse struct{}"
	}

	return strings.Join([]string{"MigrateLogicDbResponse", string(data)}, " ")
}
