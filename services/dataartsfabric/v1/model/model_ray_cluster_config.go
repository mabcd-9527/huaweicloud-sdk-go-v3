package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RayClusterConfig **参数解释**：Ray集群配置。 **约束限制**：不涉及。
type RayClusterConfig struct {

	// **参数解释**：是否开启资源自定义。 **约束限制**：不涉及。 **取值范围**：开启true，关闭false。 **默认取值**：false。
	EnableCustom *bool `json:"enable_custom,omitempty"`

	HeadGroupSpec *HeadGroupSpecV2 `json:"head_group_spec"`

	// **参数解释**：Worker Group的配置。 **约束限制**：[1,1000]。
	WorkerGroupSpecs []WorkerGroupSpecV2 `json:"worker_group_specs"`
}

func (o RayClusterConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RayClusterConfig struct{}"
	}

	return strings.Join([]string{"RayClusterConfig", string(data)}, " ")
}
