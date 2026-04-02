package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGcTaskRequest Request Object
type ShowGcTaskRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	// 任务ID
	GcId int32 `json:"gc_id"`
}

func (o ShowGcTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGcTaskRequest struct{}"
	}

	return strings.Join([]string{"ShowGcTaskRequest", string(data)}, " ")
}
