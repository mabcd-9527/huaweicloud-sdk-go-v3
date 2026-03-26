package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableAgency - **参数解释**：是否开启Ray Service委托授权功能。 - **约束限制**：不涉及。 - **取值范围**：开启true，关闭false。 - **默认取值**：false。
type EnableAgency struct {
}

func (o EnableAgency) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableAgency struct{}"
	}

	return strings.Join([]string{"EnableAgency", string(data)}, " ")
}
