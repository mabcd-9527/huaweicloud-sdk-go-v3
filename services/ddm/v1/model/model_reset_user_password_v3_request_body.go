package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetUserPasswordV3RequestBody 重置账号密码的请求对象
type ResetUserPasswordV3RequestBody struct {

	// **参数解释**：  实例账号密码。  **约束限制**：  - 长度为8~32个字符。 - 至少包含三种字符组合：小写字母、大写字母、数字、特殊字符 ~ ! @ # % ^ * - _ + ? - 不能使用简单、强度不够、容易被猜测的密码。 - 不能与账号名称或者倒序的账号名称相同。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Password string `json:"password"`
}

func (o ResetUserPasswordV3RequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetUserPasswordV3RequestBody struct{}"
	}

	return strings.Join([]string{"ResetUserPasswordV3RequestBody", string(data)}, " ")
}
