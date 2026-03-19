package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type BatchCreateIpReputationRuleRequestBody struct {

	// **参数解释：** 规则名称 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Name string `json:"name"`

	// **参数解释：** 信誉类型（目前只支持idc） **约束限制：** 不涉及 **取值范围：** - idc  **默认取值：** 不涉及
	Type BatchCreateIpReputationRuleRequestBodyType `json:"type"`

	// **参数解释：** 标签列表，用于指定关联的情报标识，可通过ConfirmPolicyIpReputationMap接口查询；多个标识之间使用逗号\",\"分隔 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Tags []string `json:"tags"`

	Action *CreateIpReputationRuleRequestBodyAction `json:"action"`

	// **参数解释：** 规则描述 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释：** 添加规则的策略id列表。策略id从\"查询防护策略列表\"(ListPolicy)接口获取，多个策略之间用“,”隔开 **约束限制：** 不能为空 **取值范围：** 不涉及 **默认取值：** 不涉及
	PolicyIds []string `json:"policy_ids"`
}

func (o BatchCreateIpReputationRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateIpReputationRuleRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreateIpReputationRuleRequestBody", string(data)}, " ")
}

type BatchCreateIpReputationRuleRequestBodyType struct {
	value string
}

type BatchCreateIpReputationRuleRequestBodyTypeEnum struct {
	IDC BatchCreateIpReputationRuleRequestBodyType
}

func GetBatchCreateIpReputationRuleRequestBodyTypeEnum() BatchCreateIpReputationRuleRequestBodyTypeEnum {
	return BatchCreateIpReputationRuleRequestBodyTypeEnum{
		IDC: BatchCreateIpReputationRuleRequestBodyType{
			value: "idc",
		},
	}
}

func (c BatchCreateIpReputationRuleRequestBodyType) Value() string {
	return c.value
}

func (c BatchCreateIpReputationRuleRequestBodyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchCreateIpReputationRuleRequestBodyType) UnmarshalJSON(b []byte) error {
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
