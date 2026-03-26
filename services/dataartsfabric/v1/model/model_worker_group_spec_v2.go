package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkerGroupSpecV2 **参数解释**：Ray worker配置信息V2。 **约束限制**：不涉及。
type WorkerGroupSpecV2 struct {

	// **参数解释**：名称。 **约束限制**：不涉及。 **取值范围**：长度[0,128]字母、数字、下划线(_)、中划线(-)的组合。 **默认取值**：不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：最小副本数。 **约束限制**：不涉及。 **取值范围**：[0,10000]。 **默认取值**：不涉及。
	MinReplicas *int32 `json:"min_replicas,omitempty"`

	// **参数解释**：最大副本数。 **约束限制**：不涉及。 **取值范围**：[1,10000]。 **默认取值**：不涉及。
	MaxReplicas *int32 `json:"max_replicas,omitempty"`

	Limits *ResourceSpec `json:"limits,omitempty"`

	Requests *ResourceSpec `json:"requests,omitempty"`
}

func (o WorkerGroupSpecV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkerGroupSpecV2 struct{}"
	}

	return strings.Join([]string{"WorkerGroupSpecV2", string(data)}, " ")
}
