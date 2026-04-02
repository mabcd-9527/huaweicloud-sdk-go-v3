package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGcScheduleRequest Request Object
type UpdateGcScheduleRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	Body *UpdateScheduleRequestBody `json:"body,omitempty"`
}

func (o UpdateGcScheduleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGcScheduleRequest struct{}"
	}

	return strings.Join([]string{"UpdateGcScheduleRequest", string(data)}, " ")
}
