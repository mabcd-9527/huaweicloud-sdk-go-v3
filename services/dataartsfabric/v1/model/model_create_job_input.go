package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateJobInput 创建的输入请求
type CreateJobInput struct {

	// Job名称,只能包含中文、字母、数字、下划线、中划线、点、空格
	Name string `json:"name"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	Type *JobType `json:"type"`

	Version *JobVersionInput `json:"version"`

	// Endpoint Id，32~36位的英文、数字、短横组合。
	CurrentEndpointId *string `json:"current_endpoint_id,omitempty"`
}

func (o CreateJobInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateJobInput struct{}"
	}

	return strings.Join([]string{"CreateJobInput", string(data)}, " ")
}
