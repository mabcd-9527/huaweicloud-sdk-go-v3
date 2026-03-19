package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateIgnoreRuleResponse Response Object
type BatchCreateIgnoreRuleResponse struct {

	// 规则id
	Id *string `json:"id,omitempty"`

	// 策略id
	Policyid *string `json:"policyid,omitempty"`

	// 创建规则的时间戳
	Timestamp *int64 `json:"timestamp,omitempty"`

	// 规则描述
	Description *string `json:"description,omitempty"`

	// **参数解释：** 规则状态标识，用于指定规则的启用或关闭状态 **约束限制：** 不涉及 **取值范围：**  - 0：关闭  - 1：开启 **默认取值：** 不涉及
	Status *int32 `json:"status,omitempty"`

	// 被屏蔽检测的规则类型或规则ID
	Rule *string `json:"rule,omitempty"`

	// 版本号固定值为1,代表v2版本误报屏蔽规则，v1版本仅支持兼容旧版本，不支持创建
	Mode *int32 `json:"mode,omitempty"`

	// 条件列表
	Conditions *[]Condition `json:"conditions,omitempty"`

	// 附加条件
	MultiCondition *bool `json:"multiCondition,omitempty"`

	// 引用表来源，1代表用户创建，其它值代表modulleX自动生成
	Producer *int32 `json:"producer,omitempty"`

	Advanced *IgnoreAdvanced `json:"advanced,omitempty"`

	// 防护域名或防护网站
	Domain         *[]string `json:"domain,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o BatchCreateIgnoreRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateIgnoreRuleResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateIgnoreRuleResponse", string(data)}, " ")
}
