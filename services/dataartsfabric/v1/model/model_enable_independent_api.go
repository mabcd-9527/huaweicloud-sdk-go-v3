package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableIndependentApi - **参数解释**：是否开启Ray Service独立Api功能。 - **约束限制**：不涉及。 - **取值范围**：开启true，关闭false。 - **默认取值**：false。
type EnableIndependentApi struct {
}

func (o EnableIndependentApi) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableIndependentApi struct{}"
	}

	return strings.Join([]string{"EnableIndependentApi", string(data)}, " ")
}
