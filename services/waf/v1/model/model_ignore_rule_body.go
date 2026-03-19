package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IgnoreRuleBody struct {

	// 规则id
	Id *string `json:"id,omitempty"`

	// 该规则属于的防护策略的id
	Policyid *string `json:"policyid,omitempty"`

	// 创建规则的时间戳
	Timestamp *int64 `json:"timestamp,omitempty"`

	// 规则描述
	Description *string `json:"description,omitempty"`

	// **参数解释：** 规则状态标识，用于指定规则的启用或关闭状态 **约束限制：** 不涉及 **取值范围：**  - 0：关闭  - 1：开启 **默认取值：** 不涉及
	Status *int32 `json:"status,omitempty"`

	// 误报规则屏蔽路径，仅在mode为0的状态下有该字段
	Url *string `json:"url,omitempty"`

	// 被屏蔽检测的规则类型或规则ID
	Rule *string `json:"rule,omitempty"`

	// 版本号，0代表v1旧版本，1代表v2新版本；mode为0时，不存在conditions和multi_conditions字段，存在url和url_logic字段；mode为1时，不存在url和url_logic字段，存在conditions或multi_conditions字段，具体以实际返回结果为准
	Mode *int32 `json:"mode,omitempty"`

	// 匹配逻辑支持（equal：等于，not_equal：不等于，contain：包含，not_contain：不包含，prefix：前缀为，not_prefix：前缀不为，suffix：后缀为，not_suffix：后缀不为，regular_match：正则匹配，regular_not_match：正则不匹配）
	UrlLogic *string `json:"url_logic,omitempty"`

	// 条件列表
	Conditions *[]Condition `json:"conditions,omitempty"`

	// 防护域名或防护网站
	Domain *[]string `json:"domain,omitempty"`

	Advanced *IgnoreAdvanced `json:"advanced,omitempty"`
}

func (o IgnoreRuleBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IgnoreRuleBody struct{}"
	}

	return strings.Join([]string{"IgnoreRuleBody", string(data)}, " ")
}
