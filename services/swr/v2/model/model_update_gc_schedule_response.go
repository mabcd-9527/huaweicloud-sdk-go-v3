package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGcScheduleResponse Response Object
type UpdateGcScheduleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateGcScheduleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGcScheduleResponse struct{}"
	}

	return strings.Join([]string{"UpdateGcScheduleResponse", string(data)}, " ")
}
