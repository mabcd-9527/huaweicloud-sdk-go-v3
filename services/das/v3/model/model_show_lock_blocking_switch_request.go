package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowLockBlockingSwitchRequest Request Object
type ShowLockBlockingSwitchRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 数据库类型。当前支持的数据库类型包括SQLServer
	EngineType string `json:"engine_type"`

	// 请求语言类型。en-us：英文。 zh-cn：中文。
	XLanguage *ShowLockBlockingSwitchRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ShowLockBlockingSwitchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLockBlockingSwitchRequest struct{}"
	}

	return strings.Join([]string{"ShowLockBlockingSwitchRequest", string(data)}, " ")
}

type ShowLockBlockingSwitchRequestXLanguage struct {
	value string
}

type ShowLockBlockingSwitchRequestXLanguageEnum struct {
	ZH_CN ShowLockBlockingSwitchRequestXLanguage
	EN_US ShowLockBlockingSwitchRequestXLanguage
}

func GetShowLockBlockingSwitchRequestXLanguageEnum() ShowLockBlockingSwitchRequestXLanguageEnum {
	return ShowLockBlockingSwitchRequestXLanguageEnum{
		ZH_CN: ShowLockBlockingSwitchRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowLockBlockingSwitchRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowLockBlockingSwitchRequestXLanguage) Value() string {
	return c.value
}

func (c ShowLockBlockingSwitchRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowLockBlockingSwitchRequestXLanguage) UnmarshalJSON(b []byte) error {
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
