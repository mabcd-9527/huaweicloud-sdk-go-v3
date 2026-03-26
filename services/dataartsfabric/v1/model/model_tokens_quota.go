package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TokensQuota **参数解释**：Tokens配额信息。 **约束限制**：不涉及。
type TokensQuota struct {

	// 总配额。
	Total *int64 `json:"total,omitempty"`

	// 已使用配额。
	Used *int64 `json:"used,omitempty"`

	// **参数解释**：到期时间。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	ExpireTime *sdktime.SdkTime `json:"expire_time,omitempty"`
}

func (o TokensQuota) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TokensQuota struct{}"
	}

	return strings.Join([]string{"TokensQuota", string(data)}, " ")
}
