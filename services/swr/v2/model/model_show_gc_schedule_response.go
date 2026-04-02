package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGcScheduleResponse Response Object
type ShowGcScheduleResponse struct {
	Schedule *ScheduleDetails `json:"schedule,omitempty"`

	Parameters     *GcParameters `json:"parameters,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowGcScheduleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGcScheduleResponse struct{}"
	}

	return strings.Join([]string{"ShowGcScheduleResponse", string(data)}, " ")
}
