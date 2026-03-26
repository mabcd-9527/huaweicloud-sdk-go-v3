package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableCustom **参数解释**：是否开启资源自定义。 **约束限制**：不涉及。 **取值范围**：开启true，关闭false。 **默认取值**：false。
type EnableCustom struct {
}

func (o EnableCustom) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableCustom struct{}"
	}

	return strings.Join([]string{"EnableCustom", string(data)}, " ")
}
