package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Deployment **参数解释**：Deployment配置信息。 **约束限制**：不涉及。
type Deployment struct {

	// **参数解释**：Deployment的名称。 **约束限制**：不涉及。 **取值范围**：长度为[1,64]的中文、字母、数字、下划线、中划线、半角句号（.）、空格的组合。 **默认取值**：不涉及。
	Name *string `json:"name,omitempty"`

	RayActorOptions *RayActorOptions `json:"ray_actor_options,omitempty"`

	AutoscalingConfig *AutoscalingConfig `json:"autoscaling_config,omitempty"`

	// **参数解释**：副本数量。 **约束限制**：不涉及。 **取值范围**：[0,1000]。 **默认取值**：不涉及。
	NumReplicas *int32 `json:"num_replicas,omitempty"`

	// **参数解释**：用户自定义配置。 **约束限制**：不涉及。
	UserConfig *interface{} `json:"user_config,omitempty"`

	// **参数解释**：每个节点允许的最大副本数。 **约束限制**：不涉及。 **取值范围**：[1,100]。 **默认取值**：1。
	MaxReplicasPerNode *int32 `json:"max_replicas_per_node,omitempty"`

	// **参数解释**：每个副本可接受的最大并发请求数。 **约束限制**：不涉及。 **取值范围**：[1,100000]。 **默认取值**：不涉及。
	MaxOngoingRequests *int32 `json:"max_ongoing_requests,omitempty"`

	// **参数解释**：deployment可接受的排队的最大请求数。-1时表示无限制。 **约束限制**：不涉及。 **取值范围**：[-1,100000]。 **默认取值**：不涉及。
	MaxQueuedRequests *int32 `json:"max_queued_requests,omitempty"`
}

func (o Deployment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Deployment struct{}"
	}

	return strings.Join([]string{"Deployment", string(data)}, " ")
}
