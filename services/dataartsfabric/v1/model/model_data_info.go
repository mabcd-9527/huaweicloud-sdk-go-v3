package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DataInfo **参数解释**：数据信息。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
type DataInfo struct {

	// **参数解释**：地址。 **约束限制**：不涉及。 **取值范围**：长度[4,1024]。 **默认取值**：不涉及。
	InputPath *string `json:"input_path,omitempty"`

	// **参数解释**：地址。 **约束限制**：不涉及。 **取值范围**：长度[4,1024]。 **默认取值**：不涉及。
	OutputPath *string `json:"output_path,omitempty"`

	// **参数解释**：环境变量名称。 **约束限制**：不涉及。 **取值范围**：长度[2,100]。 **默认取值**：不涉及。
	EnvVarName *string `json:"env_var_name,omitempty"`
}

func (o DataInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataInfo struct{}"
	}

	return strings.Join([]string{"DataInfo", string(data)}, " ")
}
