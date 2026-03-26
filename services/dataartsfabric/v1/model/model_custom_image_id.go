package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CustomImageId - **参数解释**：CustomImageId。 - **约束限制**：不涉及。 - **取值范围**：长度为[32,36]的英文字符、数字和中划线(-)的组合。 - **默认取值**：不涉及。
type CustomImageId struct {
}

func (o CustomImageId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomImageId struct{}"
	}

	return strings.Join([]string{"CustomImageId", string(data)}, " ")
}
