package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HeadGroupSpecV2 **参数解释**：HeadGroup的配置。 **约束限制**：不涉及。
type HeadGroupSpecV2 struct {
	Limits *ResourceSpec `json:"limits,omitempty"`

	Requests *ResourceSpec `json:"requests,omitempty"`
}

func (o HeadGroupSpecV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HeadGroupSpecV2 struct{}"
	}

	return strings.Join([]string{"HeadGroupSpecV2", string(data)}, " ")
}
