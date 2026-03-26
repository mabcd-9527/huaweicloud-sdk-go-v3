package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CapacityReservationSpecification struct {

	// 目标容量预留ID
	Id *string `json:"id,omitempty"`

	// 实例启动的私有池容量选项
	Preference *CapacityReservationSpecificationPreference `json:"preference,omitempty"`
}

func (o CapacityReservationSpecification) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CapacityReservationSpecification struct{}"
	}

	return strings.Join([]string{"CapacityReservationSpecification", string(data)}, " ")
}

type CapacityReservationSpecificationPreference struct {
	value string
}

type CapacityReservationSpecificationPreferenceEnum struct {
	NONE     CapacityReservationSpecificationPreference
	TARGETED CapacityReservationSpecificationPreference
}

func GetCapacityReservationSpecificationPreferenceEnum() CapacityReservationSpecificationPreferenceEnum {
	return CapacityReservationSpecificationPreferenceEnum{
		NONE: CapacityReservationSpecificationPreference{
			value: "none",
		},
		TARGETED: CapacityReservationSpecificationPreference{
			value: "targeted",
		},
	}
}

func (c CapacityReservationSpecificationPreference) Value() string {
	return c.value
}

func (c CapacityReservationSpecificationPreference) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CapacityReservationSpecificationPreference) UnmarshalJSON(b []byte) error {
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
