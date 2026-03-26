package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DataBriefInfo **参数解释**：数据路径信息。 **约束限制**：不涉及。
type DataBriefInfo struct {

	// **参数解释**：地址。 **约束限制**：不涉及。 **取值范围**：长度[4,1024]。 **默认取值**：不涉及。
	InputPath *string `json:"input_path,omitempty"`
}

func (o DataBriefInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataBriefInfo struct{}"
	}

	return strings.Join([]string{"DataBriefInfo", string(data)}, " ")
}
