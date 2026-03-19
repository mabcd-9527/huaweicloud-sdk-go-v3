package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGeoipRuleResponse Response Object
type UpdateGeoipRuleResponse struct {

	// 规则id
	Id *string `json:"id,omitempty"`

	// 地理位置控制规则名称
	Name *string `json:"name,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 策略id
	Policyid *string `json:"policyid,omitempty"`

	// 地理位置封禁区域
	Geoip *string `json:"geoip,omitempty"`

	// 防护动作：  - 0 拦截  - 1 放行  - 2 仅记录
	White          *int32 `json:"white,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o UpdateGeoipRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGeoipRuleResponse struct{}"
	}

	return strings.Join([]string{"UpdateGeoipRuleResponse", string(data)}, " ")
}
