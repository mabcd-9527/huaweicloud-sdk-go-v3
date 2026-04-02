package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExecuteScheduleRequestBody struct {
	Schedule *ExecuteScheduleObj `json:"schedule"`

	Parameters *GcParameters `json:"parameters"`
}

func (o ExecuteScheduleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteScheduleRequestBody struct{}"
	}

	return strings.Join([]string{"ExecuteScheduleRequestBody", string(data)}, " ")
}
