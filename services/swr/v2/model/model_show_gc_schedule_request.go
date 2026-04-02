package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGcScheduleRequest Request Object
type ShowGcScheduleRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`
}

func (o ShowGcScheduleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGcScheduleRequest struct{}"
	}

	return strings.Join([]string{"ShowGcScheduleRequest", string(data)}, " ")
}
