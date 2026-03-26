package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AutoscalingConfig **参数解释**：自动扩缩的配置。 **约束限制**：不涉及。
type AutoscalingConfig struct {

	// **参数解释**：每个副本单位时间能提供的平均请求数。 **约束限制**：不涉及。 **取值范围**：[1,1000]。 **默认取值**：不涉及。
	TargetOngoingRequests *int32 `json:"target_ongoing_requests,omitempty"`

	// **参数解释**：最小副本数。 **约束限制**：不涉及。 **取值范围**：[0,1000]。 **默认取值**：不涉及。
	MinReplicas *int32 `json:"min_replicas,omitempty"`

	// **参数解释**：最大副本数。 **约束限制**：不涉及。 **取值范围**：[0,1000]。 **默认取值**：不涉及。
	MaxReplicas *int32 `json:"max_replicas,omitempty"`

	// **参数解释**：初始副本数。 **约束限制**：不涉及。 **取值范围**：[0,1000]。 **默认取值**：不涉及。
	InitialReplicas *int32 `json:"initial_replicas,omitempty"`

	// **参数解释**：扩容之前的等待时间。 **约束限制**：不涉及。 **取值范围**：[0,86400]。 **默认取值**：不涉及。
	UpscaleDelayS *int32 `json:"upscale_delay_s,omitempty"`

	// **参数解释**：缩容之前的等待时间。 **约束限制**：不涉及。 **取值范围**：[0,86400]。 **默认取值**：不涉及。
	DownscaleDelayS *int32 `json:"downscale_delay_s,omitempty"`
}

func (o AutoscalingConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoscalingConfig struct{}"
	}

	return strings.Join([]string{"AutoscalingConfig", string(data)}, " ")
}
