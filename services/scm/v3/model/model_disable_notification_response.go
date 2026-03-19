package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisableNotificationResponse Response Object
type DisableNotificationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DisableNotificationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableNotificationResponse struct{}"
	}

	return strings.Join([]string{"DisableNotificationResponse", string(data)}, " ")
}
