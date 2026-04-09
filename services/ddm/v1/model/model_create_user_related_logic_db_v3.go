package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateUserRelatedLogicDbV3 创建账号时关联的逻辑库请求对象
type CreateUserRelatedLogicDbV3 struct {

	// **参数解释**：  已创建的逻辑库名称。 表示给账号授予逻辑库权限。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Name string `json:"name"`
}

func (o CreateUserRelatedLogicDbV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUserRelatedLogicDbV3 struct{}"
	}

	return strings.Join([]string{"CreateUserRelatedLogicDbV3", string(data)}, " ")
}
