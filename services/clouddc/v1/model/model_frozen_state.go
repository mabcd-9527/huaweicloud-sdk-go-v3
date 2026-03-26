package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// FrozenState **参数解释**： 服务器冻结状态 **约束限制**： 不涉及 **取值范围** - 0：未冻结。 - 1: 公安冻结。 - 2：违规冻结。 - 3：违规冻结+公安冻结。 - 4：欠费冻结。 - 5：欠费冻结+公安冻结。 - 6：欠费冻结+违规冻结。 - 7：欠费冻结+公安冻结+违规冻结。 ``` **默认取值**： 不涉及
type FrozenState struct {
	value int32
}

type FrozenStateEnum struct {
	E_0 FrozenState
	E_1 FrozenState
	E_2 FrozenState
	E_3 FrozenState
	E_4 FrozenState
	E_5 FrozenState
	E_6 FrozenState
	E_7 FrozenState
}

func GetFrozenStateEnum() FrozenStateEnum {
	return FrozenStateEnum{
		E_0: FrozenState{
			value: 0,
		}, E_1: FrozenState{
			value: 1,
		}, E_2: FrozenState{
			value: 2,
		}, E_3: FrozenState{
			value: 3,
		}, E_4: FrozenState{
			value: 4,
		}, E_5: FrozenState{
			value: 5,
		}, E_6: FrozenState{
			value: 6,
		}, E_7: FrozenState{
			value: 7,
		},
	}
}

func (c FrozenState) Value() int32 {
	return c.value
}

func (c FrozenState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FrozenState) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
