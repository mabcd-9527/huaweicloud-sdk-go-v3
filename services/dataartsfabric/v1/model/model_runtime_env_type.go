package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// RuntimeEnvType **参数解释**：运行环境。 **约束限制**：不涉及。 **取值范围**：   - RELEASE：表示运行在生产引擎实例环境中；   - GRAY：表示运行在灰度引擎实例环境中。 **默认取值**：不涉及。
type RuntimeEnvType struct {
	value string
}

type RuntimeEnvTypeEnum struct {
	RELEASE RuntimeEnvType
	GRAY    RuntimeEnvType
}

func GetRuntimeEnvTypeEnum() RuntimeEnvTypeEnum {
	return RuntimeEnvTypeEnum{
		RELEASE: RuntimeEnvType{
			value: "RELEASE",
		},
		GRAY: RuntimeEnvType{
			value: "GRAY",
		},
	}
}

func (c RuntimeEnvType) Value() string {
	return c.value
}

func (c RuntimeEnvType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RuntimeEnvType) UnmarshalJSON(b []byte) error {
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
