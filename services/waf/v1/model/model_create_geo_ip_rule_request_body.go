package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateGeoIpRuleRequestBody struct {

	// 地理位置控制规则名称
	Name *string `json:"name,omitempty"`

	// 地理位置封禁区域，选择区域对应的字母代号,用中划线|分隔，可通过ShowPolicyGeoipMap接口查询支持的区域
	Geoip string `json:"geoip"`

	// 防护动作：  - 0 拦截  - 1 放行  - 2 仅记录
	White int32 `json:"white"`

	// **参数解释：** 规则状态标识，用于指定规则的启用或关闭状态 **约束限制：** 不涉及 **取值范围：**  - 0：关闭  - 1：开启 **默认取值：** 不涉及
	Status *int32 `json:"status,omitempty"`

	// 规则描述
	Description *string `json:"description,omitempty"`
}

func (o CreateGeoIpRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGeoIpRuleRequestBody struct{}"
	}

	return strings.Join([]string{"CreateGeoIpRuleRequestBody", string(data)}, " ")
}
