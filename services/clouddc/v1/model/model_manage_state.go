package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ManageState **参数解释**： 服务器管理状态 **约束限制**： 不涉及 **取值范围** - onboard：上架中，用户下单，完成LLD设计。 - os-ready: 软调完成 服务器完成软调，即将进行转维。 - ready：交付完成，完成硬装、网调、服务器初始化、软调及转维验收。 - in-use：使用中，用户发放裸机。 - frozen：冻结，因欠费导致资源冻结。 - maintain:维护中 服务器故障，进入维修 - offboarding：下架中。  ```mermaid stateDiagram-v2    [*] --> onboard : 完成LLD设计   onboard --> ready : 完成网调、服务器初始化、软调及转维验收   ready --> in_use : 发放裸机实例   ready --> offboarding : 请求下架   ready --> frozen : 欠费      in_use --> ready : 删除裸机实例   in_use --> frozen : 欠费    frozen --> offboarding : 请求下架   in_use --> offboarding : 请求下架   offboarding --> [*] : 完成下架   state \"in-use\" as in_use ``` **默认取值**： 不涉及
type ManageState struct {
	value string
}

type ManageStateEnum struct {
	ONBOARD     ManageState
	OS_READY    ManageState
	READY       ManageState
	IN_USE      ManageState
	MAINTAIN    ManageState
	FROZEN      ManageState
	OFFBOARDING ManageState
}

func GetManageStateEnum() ManageStateEnum {
	return ManageStateEnum{
		ONBOARD: ManageState{
			value: "onboard",
		},
		OS_READY: ManageState{
			value: "os-ready",
		},
		READY: ManageState{
			value: "ready",
		},
		IN_USE: ManageState{
			value: "in-use",
		},
		MAINTAIN: ManageState{
			value: "maintain",
		},
		FROZEN: ManageState{
			value: "frozen",
		},
		OFFBOARDING: ManageState{
			value: "offboarding",
		},
	}
}

func (c ManageState) Value() string {
	return c.value
}

func (c ManageState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ManageState) UnmarshalJSON(b []byte) error {
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
