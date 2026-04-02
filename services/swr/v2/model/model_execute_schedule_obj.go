package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ExecuteScheduleObj struct {

	// 计划类型。有效值为 'Manual'（手动）。  'Manual' 表示立即触发。
	Type ExecuteScheduleObjType `json:"type"`
}

func (o ExecuteScheduleObj) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteScheduleObj struct{}"
	}

	return strings.Join([]string{"ExecuteScheduleObj", string(data)}, " ")
}

type ExecuteScheduleObjType struct {
	value string
}

type ExecuteScheduleObjTypeEnum struct {
	MANUAL ExecuteScheduleObjType
}

func GetExecuteScheduleObjTypeEnum() ExecuteScheduleObjTypeEnum {
	return ExecuteScheduleObjTypeEnum{
		MANUAL: ExecuteScheduleObjType{
			value: "Manual",
		},
	}
}

func (c ExecuteScheduleObjType) Value() string {
	return c.value
}

func (c ExecuteScheduleObjType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExecuteScheduleObjType) UnmarshalJSON(b []byte) error {
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
