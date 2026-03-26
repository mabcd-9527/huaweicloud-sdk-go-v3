package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableOidc - **参数解释**：是否开启oidc功能。 - **约束限制**：不涉及。 - **取值范围**：开启true，关闭false。 - **默认取值**：false。
type EnableOidc struct {
}

func (o EnableOidc) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableOidc struct{}"
	}

	return strings.Join([]string{"EnableOidc", string(data)}, " ")
}
