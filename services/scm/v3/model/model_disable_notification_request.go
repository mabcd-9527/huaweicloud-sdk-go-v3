package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisableNotificationRequest Request Object
type DisableNotificationRequest struct {

	// 禁用证书提醒列表。
	Body *[]NoticeRequestBody `json:"body,omitempty"`
}

func (o DisableNotificationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableNotificationRequest struct{}"
	}

	return strings.Join([]string{"DisableNotificationRequest", string(data)}, " ")
}
