package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RayServeConfig **参数解释**：RayServe配置。 **约束限制**：不涉及。
type RayServeConfig struct {

	// **参数解释**：Application。 **约束限制**：不涉及。 **取值范围**：[0,5]。 **默认取值**：不涉及。
	Applications *[]ServeApplication `json:"applications,omitempty"`
}

func (o RayServeConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RayServeConfig struct{}"
	}

	return strings.Join([]string{"RayServeConfig", string(data)}, " ")
}
