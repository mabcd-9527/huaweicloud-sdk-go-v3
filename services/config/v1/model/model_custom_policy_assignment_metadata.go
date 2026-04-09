package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CustomPolicyAssignmentMetadata 自定义规则元数据
type CustomPolicyAssignmentMetadata struct {

	// 规则描述
	Description *string `json:"description,omitempty"`

	// 触发周期
	Period *CustomPolicyAssignmentMetadataPeriod `json:"period,omitempty"`

	// 输入参数
	Parameters map[string]PolicyParameterValue `json:"parameters,omitempty"`

	PolicyFilter *PolicyFilterDefinition `json:"policy_filter,omitempty"`

	PolicyFilterV2 *PolicyFilterDefinitionV2 `json:"policy_filter_v2,omitempty"`

	// 自定义函数的urn
	FunctionUrn string `json:"function_urn"`
}

func (o CustomPolicyAssignmentMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomPolicyAssignmentMetadata struct{}"
	}

	return strings.Join([]string{"CustomPolicyAssignmentMetadata", string(data)}, " ")
}

type CustomPolicyAssignmentMetadataPeriod struct {
	value string
}

type CustomPolicyAssignmentMetadataPeriodEnum struct {
	ONE_HOUR          CustomPolicyAssignmentMetadataPeriod
	THREE_HOURS       CustomPolicyAssignmentMetadataPeriod
	SIX_HOURS         CustomPolicyAssignmentMetadataPeriod
	TWELVE_HOURS      CustomPolicyAssignmentMetadataPeriod
	TWENTY_FOUR_HOURS CustomPolicyAssignmentMetadataPeriod
}

func GetCustomPolicyAssignmentMetadataPeriodEnum() CustomPolicyAssignmentMetadataPeriodEnum {
	return CustomPolicyAssignmentMetadataPeriodEnum{
		ONE_HOUR: CustomPolicyAssignmentMetadataPeriod{
			value: "One_Hour",
		},
		THREE_HOURS: CustomPolicyAssignmentMetadataPeriod{
			value: "Three_Hours",
		},
		SIX_HOURS: CustomPolicyAssignmentMetadataPeriod{
			value: "Six_Hours",
		},
		TWELVE_HOURS: CustomPolicyAssignmentMetadataPeriod{
			value: "Twelve_Hours",
		},
		TWENTY_FOUR_HOURS: CustomPolicyAssignmentMetadataPeriod{
			value: "TwentyFour_Hours",
		},
	}
}

func (c CustomPolicyAssignmentMetadataPeriod) Value() string {
	return c.value
}

func (c CustomPolicyAssignmentMetadataPeriod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CustomPolicyAssignmentMetadataPeriod) UnmarshalJSON(b []byte) error {
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
