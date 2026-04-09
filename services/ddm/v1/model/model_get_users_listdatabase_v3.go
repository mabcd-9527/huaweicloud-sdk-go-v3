package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GetUsersListdatabaseV3 struct {

	// **参数解释**：  关联的逻辑库名称。  **取值范围**：  不涉及。
	Name *string `json:"name,omitempty"`
}

func (o GetUsersListdatabaseV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetUsersListdatabaseV3 struct{}"
	}

	return strings.Join([]string{"GetUsersListdatabaseV3", string(data)}, " ")
}
