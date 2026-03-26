package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Framework Frameworks详细信息
type Framework struct {

	// Framework名称
	Name *string `json:"name,omitempty"`

	// Framework版本
	VersionName *string `json:"version_name,omitempty"`
}

func (o Framework) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Framework struct{}"
	}

	return strings.Join([]string{"Framework", string(data)}, " ")
}
