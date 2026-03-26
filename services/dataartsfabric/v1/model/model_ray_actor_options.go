package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RayActorOptions **参数解释**：RayActor配置。 **约束限制**：不涉及。
type RayActorOptions struct {

	// **参数解释**：CPU数量。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	NumCpus *float64 `json:"num_cpus,omitempty"`

	// **参数解释**：内存数量。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Memory *float64 `json:"memory,omitempty"`

	// **参数解释**：资源配置。 **约束限制**：不涉及。 **取值范围**：长度[0,1024]。 **默认取值**：不涉及。
	Resources *string `json:"resources,omitempty"`
}

func (o RayActorOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RayActorOptions struct{}"
	}

	return strings.Join([]string{"RayActorOptions", string(data)}, " ")
}
