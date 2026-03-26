package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServeRuntimeEnv **参数解释**：Serve的运行时环境配置。 **约束限制**：不涉及。 **取值范围**：可选参数有：working_dir：代码将在其中运行的工作目录。必须是远程URI，如s3或本地路径;env_vars：要设置的环境变量。 **默认取值**：不涉及。
type ServeRuntimeEnv struct {

	// **参数解释**：代码将在其中运行的工作目录。 **约束限制**：必须是远程URI，如s3或本地路径。 **取值范围**：不涉及。 **默认取值**：不涉及。
	WorkingDir *string `json:"working_dir,omitempty"`

	// **参数解释**：要设置的环境变量。 **约束限制**：不涉及。
	EnvVars map[string]string `json:"env_vars,omitempty"`
}

func (o ServeRuntimeEnv) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServeRuntimeEnv struct{}"
	}

	return strings.Join([]string{"ServeRuntimeEnv", string(data)}, " ")
}
