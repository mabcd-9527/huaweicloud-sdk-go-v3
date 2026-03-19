package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableNotificationResponse Response Object
type EnableNotificationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o EnableNotificationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableNotificationResponse struct{}"
	}

	return strings.Join([]string{"EnableNotificationResponse", string(data)}, " ")
}
