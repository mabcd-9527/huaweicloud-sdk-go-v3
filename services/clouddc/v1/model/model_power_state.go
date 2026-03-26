package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PowerState **参数解释**： 电源状态 power_state 会根据不同的操作和事件发生转换，常见的状态转换流程如下：   - 开机流程：off -> on   - 关机流程：on -> off   - 重启流程：on -> off -> on  **约束限制**： 不涉及 **取值范围**： - on：表示节点的电源已开启，硬件处于通电状态，操作系统正在运行或者可以正常启动。这意味着节点能够执行计算任务，为上层应用提供服务。 - off：表明节点的电源已关闭，硬件停止供电，所有组件处于非工作状态，无法执行任何计算任务。  **默认取值**： 不涉及
type PowerState struct {
	value string
}

type PowerStateEnum struct {
	ON  PowerState
	OFF PowerState
}

func GetPowerStateEnum() PowerStateEnum {
	return PowerStateEnum{
		ON: PowerState{
			value: "on",
		},
		OFF: PowerState{
			value: "off",
		},
	}
}

func (c PowerState) Value() string {
	return c.value
}

func (c PowerState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PowerState) UnmarshalJSON(b []byte) error {
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
