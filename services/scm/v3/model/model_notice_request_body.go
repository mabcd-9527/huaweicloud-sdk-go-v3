package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NoticeRequestBody struct {

	// 资源ID。
	ResourceId *string `json:"resource_id,omitempty"`

	// 提醒类型，取值如下： - CERT_EXPIRING(证书到期) - CERT_EXPIRED(证书过期)
	NoticeType *[]string `json:"notice_type,omitempty"`
}

func (o NoticeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NoticeRequestBody struct{}"
	}

	return strings.Join([]string{"NoticeRequestBody", string(data)}, " ")
}
