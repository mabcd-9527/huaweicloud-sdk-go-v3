package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ReadOnlySwitchReq struct {

	// **参数解释**：  实例是否设置为只读。 - true：设置为只读。 - false：解除只读状态。  **约束限制**：  不涉及。  **取值范围**：  true或者false。  **默认取值**：  不涉及。
	Readonly bool `json:"readonly"`
}

func (o ReadOnlySwitchReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReadOnlySwitchReq struct{}"
	}

	return strings.Join([]string{"ReadOnlySwitchReq", string(data)}, " ")
}
