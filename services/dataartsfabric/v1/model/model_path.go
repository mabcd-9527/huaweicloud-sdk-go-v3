package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Path **参数解释**：地址。 **约束限制**：不涉及。 **取值范围**：长度[4,1024]。 **默认取值**：不涉及。
type Path struct {
}

func (o Path) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Path struct{}"
	}

	return strings.Join([]string{"Path", string(data)}, " ")
}
