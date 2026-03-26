package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClientId - **参数解释**：客户端ID。 - **约束限制**：不涉及。 - **取值范围**：长度为[1,64]的字母、数字、下划线(_)、中划线(-)、点(.)的组合。 - **默认取值**：null。
type ClientId struct {
}

func (o ClientId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClientId struct{}"
	}

	return strings.Join([]string{"ClientId", string(data)}, " ")
}
