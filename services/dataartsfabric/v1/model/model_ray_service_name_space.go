package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RayServiceNameSpace - **参数解释**：RayService命名空间。 - **约束限制**：不涉及。 - **取值范围**：长度为16位，均为ray-svc-xxx格式。 - **默认取值**：null。
type RayServiceNameSpace struct {
}

func (o RayServiceNameSpace) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RayServiceNameSpace struct{}"
	}

	return strings.Join([]string{"RayServiceNameSpace", string(data)}, " ")
}
