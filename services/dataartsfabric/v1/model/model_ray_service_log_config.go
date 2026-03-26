package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RayServiceLogConfig **参数解释**：RayService日志配置。 **约束限制**：不涉及。
type RayServiceLogConfig struct {
	Lts *RayServiceLogLtsConfig `json:"lts,omitempty"`
}

func (o RayServiceLogConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RayServiceLogConfig struct{}"
	}

	return strings.Join([]string{"RayServiceLogConfig", string(data)}, " ")
}
