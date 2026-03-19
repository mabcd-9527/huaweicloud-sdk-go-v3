package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdatePunishmentRuleRequestBody struct {

	// **参数解释：** 攻击惩罚类别 **约束限制：** 不支持修改 **取值范围：**  - long_ip_block  - long_cookie_block  - long_params_block  - long_header_block  - short_ip_block  - short_cookie_block  - short_params_block  - short_header_block **默认取值：** 不涉及
	Category UpdatePunishmentRuleRequestBodyCategory `json:"category"`

	// **参数解释：** 时间单位，会影响“拦截时间”参数的取值范围 **约束限制：** 不涉及 **取值范围：**  - SECOND  - MINUTE  - HOUR  - DAY  - MONTH **默认取值：** SECOND
	TimeUnit *UpdatePunishmentRuleRequestBodyTimeUnit `json:"time_unit,omitempty"`

	// **参数解释：** 拦截时间 **约束限制：** 取值范围取决于惩罚类别和时间单位 **取值范围：** - short_xxx   - SECOND  [1, 300] - long_xxx   - SECOND [301, 7776000]   - MINUTE [6, 129600]   - HOUR [1, 2160]   - DAY [1, 90]   - MONTH [1, 3] **默认取值：** 不涉及
	BlockTime int32 `json:"block_time"`

	// 规则描述
	Description *string `json:"description,omitempty"`
}

func (o UpdatePunishmentRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePunishmentRuleRequestBody struct{}"
	}

	return strings.Join([]string{"UpdatePunishmentRuleRequestBody", string(data)}, " ")
}

type UpdatePunishmentRuleRequestBodyCategory struct {
	value string
}

type UpdatePunishmentRuleRequestBodyCategoryEnum struct {
	LONG_IP_BLOCK      UpdatePunishmentRuleRequestBodyCategory
	LONG_COOKIE_BLOCK  UpdatePunishmentRuleRequestBodyCategory
	LONG_PARAMS_BLOCK  UpdatePunishmentRuleRequestBodyCategory
	LONG_HEADER_BLOCK  UpdatePunishmentRuleRequestBodyCategory
	SHORT_IP_BLOCK     UpdatePunishmentRuleRequestBodyCategory
	SHORT_COOKIE_BLOCK UpdatePunishmentRuleRequestBodyCategory
	SHORT_PARAMS_BLOCK UpdatePunishmentRuleRequestBodyCategory
	SHORT_HEADER_BLOCK UpdatePunishmentRuleRequestBodyCategory
}

func GetUpdatePunishmentRuleRequestBodyCategoryEnum() UpdatePunishmentRuleRequestBodyCategoryEnum {
	return UpdatePunishmentRuleRequestBodyCategoryEnum{
		LONG_IP_BLOCK: UpdatePunishmentRuleRequestBodyCategory{
			value: "long_ip_block",
		},
		LONG_COOKIE_BLOCK: UpdatePunishmentRuleRequestBodyCategory{
			value: "long_cookie_block",
		},
		LONG_PARAMS_BLOCK: UpdatePunishmentRuleRequestBodyCategory{
			value: "long_params_block",
		},
		LONG_HEADER_BLOCK: UpdatePunishmentRuleRequestBodyCategory{
			value: "long_header_block",
		},
		SHORT_IP_BLOCK: UpdatePunishmentRuleRequestBodyCategory{
			value: "short_ip_block",
		},
		SHORT_COOKIE_BLOCK: UpdatePunishmentRuleRequestBodyCategory{
			value: "short_cookie_block",
		},
		SHORT_PARAMS_BLOCK: UpdatePunishmentRuleRequestBodyCategory{
			value: "short_params_block",
		},
		SHORT_HEADER_BLOCK: UpdatePunishmentRuleRequestBodyCategory{
			value: "short_header_block",
		},
	}
}

func (c UpdatePunishmentRuleRequestBodyCategory) Value() string {
	return c.value
}

func (c UpdatePunishmentRuleRequestBodyCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdatePunishmentRuleRequestBodyCategory) UnmarshalJSON(b []byte) error {
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

type UpdatePunishmentRuleRequestBodyTimeUnit struct {
	value string
}

type UpdatePunishmentRuleRequestBodyTimeUnitEnum struct {
	SECOND UpdatePunishmentRuleRequestBodyTimeUnit
	MINUTE UpdatePunishmentRuleRequestBodyTimeUnit
	HOUR   UpdatePunishmentRuleRequestBodyTimeUnit
	DAY    UpdatePunishmentRuleRequestBodyTimeUnit
	MONTH  UpdatePunishmentRuleRequestBodyTimeUnit
}

func GetUpdatePunishmentRuleRequestBodyTimeUnitEnum() UpdatePunishmentRuleRequestBodyTimeUnitEnum {
	return UpdatePunishmentRuleRequestBodyTimeUnitEnum{
		SECOND: UpdatePunishmentRuleRequestBodyTimeUnit{
			value: "SECOND",
		},
		MINUTE: UpdatePunishmentRuleRequestBodyTimeUnit{
			value: "MINUTE",
		},
		HOUR: UpdatePunishmentRuleRequestBodyTimeUnit{
			value: "HOUR",
		},
		DAY: UpdatePunishmentRuleRequestBodyTimeUnit{
			value: "DAY",
		},
		MONTH: UpdatePunishmentRuleRequestBodyTimeUnit{
			value: "MONTH",
		},
	}
}

func (c UpdatePunishmentRuleRequestBodyTimeUnit) Value() string {
	return c.value
}

func (c UpdatePunishmentRuleRequestBodyTimeUnit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdatePunishmentRuleRequestBodyTimeUnit) UnmarshalJSON(b []byte) error {
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
