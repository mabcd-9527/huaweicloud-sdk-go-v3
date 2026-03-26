package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AgencyName - **参数解释**：授予Ray Service的委托名称。 - **约束限制**：不超过64位。 - **取值范围**：不涉及。 - **默认取值**：null。
type AgencyName struct {
}

func (o AgencyName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgencyName struct{}"
	}

	return strings.Join([]string{"AgencyName", string(data)}, " ")
}
