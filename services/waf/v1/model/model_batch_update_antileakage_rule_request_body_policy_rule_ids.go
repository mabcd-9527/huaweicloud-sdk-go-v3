package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchUpdateAntileakageRuleRequestBodyPolicyRuleIds struct {

	// **参数解释：** 策略id，唯一标识一条防护策略.策略id从\"查询防护策略列表\"(ListPolicy)接口获取 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	PolicyId *string `json:"policy_id,omitempty"`

	// **参数解释：** 规则id数组，包含当前防护策略下的单个规则ID，通过ListAntileakageRules接口查询防泄漏规则ID **约束限制：** 单条规则ID **取值范围：** 不涉及 **默认取值：** 不涉及
	RuleIds *[]string `json:"rule_ids,omitempty"`
}

func (o BatchUpdateAntileakageRuleRequestBodyPolicyRuleIds) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateAntileakageRuleRequestBodyPolicyRuleIds struct{}"
	}

	return strings.Join([]string{"BatchUpdateAntileakageRuleRequestBodyPolicyRuleIds", string(data)}, " ")
}
