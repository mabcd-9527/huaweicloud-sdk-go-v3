package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateIpReputationRuleRequestBodyAction **参数解释：** 防护动作配置 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type CreateIpReputationRuleRequestBodyAction struct {

	// **参数解释：** 动作类型 **约束限制：** 不涉及 **取值范围：** - pass :放行 - log ： 仅记录 - block： 拦截 **默认取值：** 不涉及
	Category CreateIpReputationRuleRequestBodyActionCategory `json:"category"`
}

func (o CreateIpReputationRuleRequestBodyAction) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIpReputationRuleRequestBodyAction struct{}"
	}

	return strings.Join([]string{"CreateIpReputationRuleRequestBodyAction", string(data)}, " ")
}

type CreateIpReputationRuleRequestBodyActionCategory struct {
	value string
}

type CreateIpReputationRuleRequestBodyActionCategoryEnum struct {
	LOG   CreateIpReputationRuleRequestBodyActionCategory
	PASS  CreateIpReputationRuleRequestBodyActionCategory
	BLOCK CreateIpReputationRuleRequestBodyActionCategory
}

func GetCreateIpReputationRuleRequestBodyActionCategoryEnum() CreateIpReputationRuleRequestBodyActionCategoryEnum {
	return CreateIpReputationRuleRequestBodyActionCategoryEnum{
		LOG: CreateIpReputationRuleRequestBodyActionCategory{
			value: "log",
		},
		PASS: CreateIpReputationRuleRequestBodyActionCategory{
			value: "pass",
		},
		BLOCK: CreateIpReputationRuleRequestBodyActionCategory{
			value: "block",
		},
	}
}

func (c CreateIpReputationRuleRequestBodyActionCategory) Value() string {
	return c.value
}

func (c CreateIpReputationRuleRequestBodyActionCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateIpReputationRuleRequestBodyActionCategory) UnmarshalJSON(b []byte) error {
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
