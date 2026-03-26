package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateJobInput 更新的输入请求
type UpdateJobInput struct {

	// Job名称,只能包含中文、字母、数字、下划线、中划线、点、空格
	Name *string `json:"name,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 作业版本Id
	CurrentVersionId *string `json:"current_version_id,omitempty"`

	// Endpoint Id，32~36位的英文、数字、短横组合。
	CurrentEndpointId *string `json:"current_endpoint_id,omitempty"`

	Version *JobVersionInput `json:"version,omitempty"`
}

func (o UpdateJobInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateJobInput struct{}"
	}

	return strings.Join([]string{"UpdateJobInput", string(data)}, " ")
}
