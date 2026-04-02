package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopGcTaskRequest Request Object
type StopGcTaskRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	// 任务ID
	GcId int32 `json:"gc_id"`
}

func (o StopGcTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopGcTaskRequest struct{}"
	}

	return strings.Join([]string{"StopGcTaskRequest", string(data)}, " ")
}
