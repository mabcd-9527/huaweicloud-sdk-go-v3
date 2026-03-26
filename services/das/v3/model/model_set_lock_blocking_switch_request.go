package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SetLockBlockingSwitchRequest Request Object
type SetLockBlockingSwitchRequest struct {

	// 请求语言类型。en-us：英文。 zh-cn：中文。
	XLanguage *SetLockBlockingSwitchRequestXLanguage `json:"X-Language,omitempty"`

	Body *SetLockBlockingSwitchReq `json:"body,omitempty"`
}

func (o SetLockBlockingSwitchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetLockBlockingSwitchRequest struct{}"
	}

	return strings.Join([]string{"SetLockBlockingSwitchRequest", string(data)}, " ")
}

type SetLockBlockingSwitchRequestXLanguage struct {
	value string
}

type SetLockBlockingSwitchRequestXLanguageEnum struct {
	ZH_CN SetLockBlockingSwitchRequestXLanguage
	EN_US SetLockBlockingSwitchRequestXLanguage
}

func GetSetLockBlockingSwitchRequestXLanguageEnum() SetLockBlockingSwitchRequestXLanguageEnum {
	return SetLockBlockingSwitchRequestXLanguageEnum{
		ZH_CN: SetLockBlockingSwitchRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: SetLockBlockingSwitchRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c SetLockBlockingSwitchRequestXLanguage) Value() string {
	return c.value
}

func (c SetLockBlockingSwitchRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SetLockBlockingSwitchRequestXLanguage) UnmarshalJSON(b []byte) error {
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
