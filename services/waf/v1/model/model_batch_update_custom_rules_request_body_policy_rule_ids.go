package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchUpdateCustomRulesRequestBodyPolicyRuleIds struct {

	// **参数解释：** 策略id，唯一标识一条防护策略.策略id从\"查询防护策略列表\"(ListPolicy)接口获取 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	PolicyId string `json:"policy_id"`

	// **参数解释：** 规则id数组，包含当前防护策略下的单个精准防护规则id，您可以通过调用查询精准防护规则列表（ListCustomRules）获取规则id **约束限制：** 单条规则ID **取值范围：** 不涉及 **默认取值：** 不涉及
	RuleIds []string `json:"rule_ids"`
}

func (o BatchUpdateCustomRulesRequestBodyPolicyRuleIds) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateCustomRulesRequestBodyPolicyRuleIds struct{}"
	}

	return strings.Join([]string{"BatchUpdateCustomRulesRequestBodyPolicyRuleIds", string(data)}, " ")
}
