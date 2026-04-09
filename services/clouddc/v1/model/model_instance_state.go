package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// InstanceState **参数解释**： 实例状态 **约束限制**： 不涉及 **取值范围**： - pending：实例正在启动（分配资源/启动操作系统） - running：实例正常运行（可接受SSH/RDP连接） - stopped： 实例已完全关闭（存储卷保留） - shutting-down：实例正在终止（删除流程中） - terminated：实例已终止（资源完全删除，不可恢复） - error：实例处于异常状态（资源未完全删除）  ```mermaid stateDiagram-v2     [*] --> pending: CreateInstance     pending --> running: CreateInstance Succeed     pending --> shutting_down: DeleteInstance     running --> stopped: PowerOff/PowerReboot     stopped --> running: PowerOn/Provision Succeed     stopped --> stopped: ChangePassword     pending --> error: CreateInstance Failed     stopped --> error: Provision Failed     error --> running: Retry Provision Succeed     stopped --> shutting_down: DeleteInstance     error --> shutting_down: Retry DeleteInstance     shutting_down --> error: DeleteInstance Failed     shutting_down --> terminated: DeleteInstance Succeed ```  **默认取值**： 不涉及
type InstanceState struct {
	value string
}

type InstanceStateEnum struct {
	PENDING       InstanceState
	RUNNING       InstanceState
	STOPPED       InstanceState
	SHUTTING_DOWN InstanceState
	TERMINATED    InstanceState
	ERROR         InstanceState
}

func GetInstanceStateEnum() InstanceStateEnum {
	return InstanceStateEnum{
		PENDING: InstanceState{
			value: "pending",
		},
		RUNNING: InstanceState{
			value: "running",
		},
		STOPPED: InstanceState{
			value: "stopped",
		},
		SHUTTING_DOWN: InstanceState{
			value: "shutting-down",
		},
		TERMINATED: InstanceState{
			value: "terminated",
		},
		ERROR: InstanceState{
			value: "error",
		},
	}
}

func (c InstanceState) Value() string {
	return c.value
}

func (c InstanceState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstanceState) UnmarshalJSON(b []byte) error {
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
