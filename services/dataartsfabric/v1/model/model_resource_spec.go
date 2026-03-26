package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceSpec **参数解释**：资源配置。 **约束限制**：不涉及。
type ResourceSpec struct {

	// 资源规格，从规格列表查询获取。
	SpecCode *string `json:"spec_code,omitempty"`

	// **参数解释**：cpu核数量。 **约束限制**：不涉及。 **取值范围**：[0,192]。 **默认取值**：null。
	CpuResource *int32 `json:"cpu_resource,omitempty"`

	// **参数解释**：内存大小。 **约束限制**：不涉及。 **取值范围**：[0,1536]。 **默认取值**：null。
	MemoryResource *int32 `json:"memory_resource,omitempty"`

	// **参数解释**：昇腾卡数量。 **约束限制**：不涉及。 **取值范围**：[0,8]。 **默认取值**：null。
	NpuResource *int32 `json:"npu_resource,omitempty"`
}

func (o ResourceSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceSpec struct{}"
	}

	return strings.Join([]string{"ResourceSpec", string(data)}, " ")
}
