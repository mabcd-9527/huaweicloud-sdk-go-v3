package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type BatchUpdateIpReputationRuleRequestBody struct {

	// **参数解释：** 规则名称 **约束限制：** 长度范围：[1, 256] **取值范围：** 不涉及 **默认取值：** 不涉及
	Name string `json:"name"`

	// **参数解释：** 规则描述 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Description *string `json:"description,omitempty"`

	Action *CreateIpReputationRuleRequestBodyAction `json:"action"`

	// **参数解释：** 信誉类型（目前只支持idc） **约束限制：** 不涉及 **取值范围：** - idc  **默认取值：** 不涉及
	Type BatchUpdateIpReputationRuleRequestBodyType `json:"type"`

	// **参数解释：** 标签列表，用于指定关联的情报标识，可通过ConfirmPolicyIpReputationMap接口查询；多个标识之间使用逗号\",\"分隔 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Tags []string `json:"tags"`

	// **参数解释：** 策略和规则id数组，关联防护策略与对应的规则集合 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	PolicyRuleIds []BatchUpdateIpReputationRuleRequestBodyPolicyRuleIds `json:"policy_rule_ids"`
}

func (o BatchUpdateIpReputationRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateIpReputationRuleRequestBody struct{}"
	}

	return strings.Join([]string{"BatchUpdateIpReputationRuleRequestBody", string(data)}, " ")
}

type BatchUpdateIpReputationRuleRequestBodyType struct {
	value string
}

type BatchUpdateIpReputationRuleRequestBodyTypeEnum struct {
	IDC BatchUpdateIpReputationRuleRequestBodyType
}

func GetBatchUpdateIpReputationRuleRequestBodyTypeEnum() BatchUpdateIpReputationRuleRequestBodyTypeEnum {
	return BatchUpdateIpReputationRuleRequestBodyTypeEnum{
		IDC: BatchUpdateIpReputationRuleRequestBodyType{
			value: "idc",
		},
	}
}

func (c BatchUpdateIpReputationRuleRequestBodyType) Value() string {
	return c.value
}

func (c BatchUpdateIpReputationRuleRequestBodyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchUpdateIpReputationRuleRequestBodyType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
