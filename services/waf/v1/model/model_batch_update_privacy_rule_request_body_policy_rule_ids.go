package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchUpdatePrivacyRuleRequestBodyPolicyRuleIds struct {

	// **参数解释：** 策略id，唯一标识一条防护策略.策略id从\"查询防护策略列表\"(ListPolicy)接口获取 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	PolicyId string `json:"policy_id"`

	// **参数解释：** 规则id数组，包含当前防护策略下的单个隐私屏蔽规则id，您可以通过调用查询隐私屏蔽规则列表（ListPrivacyRule）获取规则id **约束限制：** 单条规则ID **取值范围：** 不涉及 **默认取值：** 不涉及
	RuleIds []string `json:"rule_ids"`
}

func (o BatchUpdatePrivacyRuleRequestBodyPolicyRuleIds) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdatePrivacyRuleRequestBodyPolicyRuleIds struct{}"
	}

	return strings.Join([]string{"BatchUpdatePrivacyRuleRequestBodyPolicyRuleIds", string(data)}, " ")
}
