package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConsoleConfigRequest Request Object
type ShowConsoleConfigRequest struct {

	// **参数解释：** 局点ID，未携带时仅查询基础数据 **约束限制：** 华为云支持的局点ID **取值范围：** 不涉及 **默认取值：** 不涉及
	Region *string `json:"region,omitempty"`
}

func (o ShowConsoleConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConsoleConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowConsoleConfigRequest", string(data)}, " ")
}
