package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModelList **参数解释**： 模型列表 **取值范围**： 不涉及
type ModelList struct {

	// **参数解释**： 模型数量 **取值范围**： 不涉及
	TotalSize *int32 `json:"totalSize,omitempty"`

	// **参数解释**： 模型列表 **取值范围**： 不涉及
	Models *[]Model `json:"models,omitempty"`
}

func (o ModelList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModelList struct{}"
	}

	return strings.Join([]string{"ModelList", string(data)}, " ")
}
