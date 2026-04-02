package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteGcScheduleResponse Response Object
type ExecuteGcScheduleResponse struct {

	// 制品清理的任务ID
	Id             *int32 `json:"id,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ExecuteGcScheduleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteGcScheduleResponse struct{}"
	}

	return strings.Join([]string{"ExecuteGcScheduleResponse", string(data)}, " ")
}
