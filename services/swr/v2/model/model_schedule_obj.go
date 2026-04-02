package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ScheduleObj struct {

	// 计划类型。有效值为'None'（无）和 'Custom'（自定义）。  Custom' 表示按照指定的 cron 计划触发，而 'None' 则表示取消定时计划。
	Type ScheduleObjType `json:"type"`

	// cron表达式，一种基于时间的任务调度器，type设置为Custom时，需要配置此值。
	Cron *string `json:"cron,omitempty"`
}

func (o ScheduleObj) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScheduleObj struct{}"
	}

	return strings.Join([]string{"ScheduleObj", string(data)}, " ")
}

type ScheduleObjType struct {
	value string
}

type ScheduleObjTypeEnum struct {
	NONE   ScheduleObjType
	CUSTOM ScheduleObjType
}

func GetScheduleObjTypeEnum() ScheduleObjTypeEnum {
	return ScheduleObjTypeEnum{
		NONE: ScheduleObjType{
			value: "None",
		},
		CUSTOM: ScheduleObjType{
			value: "Custom",
		},
	}
}

func (c ScheduleObjType) Value() string {
	return c.value
}

func (c ScheduleObjType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScheduleObjType) UnmarshalJSON(b []byte) error {
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
