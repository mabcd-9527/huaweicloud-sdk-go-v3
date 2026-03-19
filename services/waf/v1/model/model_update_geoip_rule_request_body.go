package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateGeoipRuleRequestBody 地理位置封禁请求体
type UpdateGeoipRuleRequestBody struct {

	// 地理位置控制规则名称
	Name *string `json:"name,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 地理位置封禁区域，选择区域对应的字母代号，可通过ShowPolicyGeoipMap接口查询支持的区域
	Geoip string `json:"geoip"`

	// 防护动作：  - 0 拦截  - 1 放行  - 2 仅记录
	White int32 `json:"white"`
}

func (o UpdateGeoipRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateGeoipRuleRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateGeoipRuleRequestBody", string(data)}, " ")
}
