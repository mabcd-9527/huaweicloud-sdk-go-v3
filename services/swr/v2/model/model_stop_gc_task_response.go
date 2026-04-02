package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopGcTaskResponse Response Object
type StopGcTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StopGcTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopGcTaskResponse struct{}"
	}

	return strings.Join([]string{"StopGcTaskResponse", string(data)}, " ")
}
