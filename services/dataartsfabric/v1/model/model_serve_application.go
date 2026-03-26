package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServeApplication **参数解释**：Ray Serve中的应用。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
type ServeApplication struct {

	// **参数解释**：Application的名称。 **约束限制**：不涉及。 **取值范围**：长度为[1,64]的中文、字母、数字、下划线、中划线、半角句号（.）、空格的组合。 **默认取值**：不涉及。
	Name string `json:"name"`

	// **参数解释**：输入路径。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	ImportPath string `json:"import_path"`

	// **参数解释**：Deployment列表。 **约束限制**：不涉及。 **取值范围**：[1,5]。 **默认取值**：不涉及。
	Deployments []Deployment `json:"deployments"`

	RuntimeEnv *ServeRuntimeEnv `json:"runtime_env"`

	// **参数解释**：路由前缀。 **约束限制**：要求在RayServe中具有唯一性，不能跟其他的Application的前缀重复。 **取值范围**：不涉及。 **默认取值**：不涉及。
	RoutePrefix *string `json:"route_prefix,omitempty"`
}

func (o ServeApplication) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServeApplication struct{}"
	}

	return strings.Join([]string{"ServeApplication", string(data)}, " ")
}
