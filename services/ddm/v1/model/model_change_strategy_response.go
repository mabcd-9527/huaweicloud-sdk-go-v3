package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeStrategyResponse Response Object
type ChangeStrategyResponse struct {

	// **参数解释**：  工作流ID。  **取值范围**：  不涉及。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeStrategyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeStrategyResponse struct{}"
	}

	return strings.Join([]string{"ChangeStrategyResponse", string(data)}, " ")
}
