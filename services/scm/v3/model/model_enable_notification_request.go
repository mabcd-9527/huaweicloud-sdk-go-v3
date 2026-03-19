package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableNotificationRequest Request Object
type EnableNotificationRequest struct {

	// 启用证书提醒列表。
	Body *[]NoticeRequestBody `json:"body,omitempty"`
}

func (o EnableNotificationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableNotificationRequest struct{}"
	}

	return strings.Join([]string{"EnableNotificationRequest", string(data)}, " ")
}
