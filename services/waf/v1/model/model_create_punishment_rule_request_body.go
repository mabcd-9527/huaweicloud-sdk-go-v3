package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreatePunishmentRuleRequestBody struct {

	// **参数解释：** 攻击惩罚类别 **约束限制：** 创建后不可修改，单个类别只能存在一个规则 **取值范围：**  - long_ip_block  - long_cookie_block  - long_params_block  - long_header_block  - short_ip_block  - short_cookie_block  - short_params_block  - short_header_block **默认取值：** 不涉及
	Category CreatePunishmentRuleRequestBodyCategory `json:"category"`

	// **参数解释：** 时间单位，会影响“拦截时间”参数的取值范围 **约束限制：** 不涉及 **取值范围：**  - SECOND  - MINUTE  - HOUR  - DAY  - MONTH **默认取值：** SECOND
	TimeUnit *CreatePunishmentRuleRequestBodyTimeUnit `json:"time_unit,omitempty"`

	// **参数解释：** 拦截时间 **约束限制：** 取值范围取决于惩罚类别和时间单位 **取值范围：** - short_xxx   - SECOND  [1, 300] - long_xxx   - SECOND [301, 7776000]   - MINUTE [6, 129600]   - HOUR [1, 2160]   - DAY [1, 90]   - MONTH [1, 3] **默认取值：** 不涉及
	BlockTime int32 `json:"block_time"`

	// 规则描述
	Description *string `json:"description,omitempty"`
}

func (o CreatePunishmentRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePunishmentRuleRequestBody struct{}"
	}

	return strings.Join([]string{"CreatePunishmentRuleRequestBody", string(data)}, " ")
}

type CreatePunishmentRuleRequestBodyCategory struct {
	value string
}

type CreatePunishmentRuleRequestBodyCategoryEnum struct {
	LONG_IP_BLOCK      CreatePunishmentRuleRequestBodyCategory
	LONG_COOKIE_BLOCK  CreatePunishmentRuleRequestBodyCategory
	LONG_PARAMS_BLOCK  CreatePunishmentRuleRequestBodyCategory
	LONG_HEADER_BLOCK  CreatePunishmentRuleRequestBodyCategory
	SHORT_IP_BLOCK     CreatePunishmentRuleRequestBodyCategory
	SHORT_COOKIE_BLOCK CreatePunishmentRuleRequestBodyCategory
	SHORT_PARAMS_BLOCK CreatePunishmentRuleRequestBodyCategory
	SHORT_HEADER_BLOCK CreatePunishmentRuleRequestBodyCategory
}

func GetCreatePunishmentRuleRequestBodyCategoryEnum() CreatePunishmentRuleRequestBodyCategoryEnum {
	return CreatePunishmentRuleRequestBodyCategoryEnum{
		LONG_IP_BLOCK: CreatePunishmentRuleRequestBodyCategory{
			value: "long_ip_block",
		},
		LONG_COOKIE_BLOCK: CreatePunishmentRuleRequestBodyCategory{
			value: "long_cookie_block",
		},
		LONG_PARAMS_BLOCK: CreatePunishmentRuleRequestBodyCategory{
			value: "long_params_block",
		},
		LONG_HEADER_BLOCK: CreatePunishmentRuleRequestBodyCategory{
			value: "long_header_block",
		},
		SHORT_IP_BLOCK: CreatePunishmentRuleRequestBodyCategory{
			value: "short_ip_block",
		},
		SHORT_COOKIE_BLOCK: CreatePunishmentRuleRequestBodyCategory{
			value: "short_cookie_block",
		},
		SHORT_PARAMS_BLOCK: CreatePunishmentRuleRequestBodyCategory{
			value: "short_params_block",
		},
		SHORT_HEADER_BLOCK: CreatePunishmentRuleRequestBodyCategory{
			value: "short_header_block",
		},
	}
}

func (c CreatePunishmentRuleRequestBodyCategory) Value() string {
	return c.value
}

func (c CreatePunishmentRuleRequestBodyCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreatePunishmentRuleRequestBodyCategory) UnmarshalJSON(b []byte) error {
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

type CreatePunishmentRuleRequestBodyTimeUnit struct {
	value string
}

type CreatePunishmentRuleRequestBodyTimeUnitEnum struct {
	SECOND CreatePunishmentRuleRequestBodyTimeUnit
	MINUTE CreatePunishmentRuleRequestBodyTimeUnit
	HOUR   CreatePunishmentRuleRequestBodyTimeUnit
	DAY    CreatePunishmentRuleRequestBodyTimeUnit
	MONTH  CreatePunishmentRuleRequestBodyTimeUnit
}

func GetCreatePunishmentRuleRequestBodyTimeUnitEnum() CreatePunishmentRuleRequestBodyTimeUnitEnum {
	return CreatePunishmentRuleRequestBodyTimeUnitEnum{
		SECOND: CreatePunishmentRuleRequestBodyTimeUnit{
			value: "SECOND",
		},
		MINUTE: CreatePunishmentRuleRequestBodyTimeUnit{
			value: "MINUTE",
		},
		HOUR: CreatePunishmentRuleRequestBodyTimeUnit{
			value: "HOUR",
		},
		DAY: CreatePunishmentRuleRequestBodyTimeUnit{
			value: "DAY",
		},
		MONTH: CreatePunishmentRuleRequestBodyTimeUnit{
			value: "MONTH",
		},
	}
}

func (c CreatePunishmentRuleRequestBodyTimeUnit) Value() string {
	return c.value
}

func (c CreatePunishmentRuleRequestBodyTimeUnit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreatePunishmentRuleRequestBodyTimeUnit) UnmarshalJSON(b []byte) error {
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
