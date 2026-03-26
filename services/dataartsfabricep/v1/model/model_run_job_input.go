package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunJobInput 运行作业的输入请求
type RunJobInput struct {
	Job *JobRef `json:"job"`

	// 作业运行的名称，只能包含中文、字母、数字、下划线、中划线、点、空格
	Name string `json:"name"`

	// endpoint空间id
	EndpointId string `json:"endpoint_id"`
}

func (o RunJobInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunJobInput struct{}"
	}

	return strings.Join([]string{"RunJobInput", string(data)}, " ")
}
