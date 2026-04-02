package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteGcScheduleRequest Request Object
type ExecuteGcScheduleRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	Body *ExecuteScheduleRequestBody `json:"body,omitempty"`
}

func (o ExecuteGcScheduleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteGcScheduleRequest struct{}"
	}

	return strings.Join([]string{"ExecuteGcScheduleRequest", string(data)}, " ")
}
