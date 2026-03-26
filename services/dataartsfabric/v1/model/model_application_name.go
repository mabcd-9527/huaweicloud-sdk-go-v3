package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApplicationName **参数解释**：Application的名称。 **约束限制**：不涉及。 **取值范围**：长度为[1,64]的中文、字母、数字、下划线、中划线、半角句号（.）、空格的组合。 **默认取值**：不涉及。
type ApplicationName struct {
}

func (o ApplicationName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplicationName struct{}"
	}

	return strings.Join([]string{"ApplicationName", string(data)}, " ")
}
