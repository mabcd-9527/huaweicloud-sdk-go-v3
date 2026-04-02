package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ModifySecurityGroupRequest Request Object
type ModifySecurityGroupRequest struct {

	// **参数解释**: 语言。 **约束限制**: 不涉及。 **取值范围**: - zh-cn  - en-us  **默认取值**: en-us
	XLanguage *ModifySecurityGroupRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**: 实例ID，此参数是用户创建实例的唯一标识。 **约束限制**: 不涉及。 **取值范围**: 只能由英文字母、数字组成，且长度为36个字符。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	Body *ModifySecurityGroupRequestBody `json:"body,omitempty"`
}

func (o ModifySecurityGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifySecurityGroupRequest struct{}"
	}

	return strings.Join([]string{"ModifySecurityGroupRequest", string(data)}, " ")
}

type ModifySecurityGroupRequestXLanguage struct {
	value string
}

type ModifySecurityGroupRequestXLanguageEnum struct {
	ZH_CN ModifySecurityGroupRequestXLanguage
	EN_US ModifySecurityGroupRequestXLanguage
}

func GetModifySecurityGroupRequestXLanguageEnum() ModifySecurityGroupRequestXLanguageEnum {
	return ModifySecurityGroupRequestXLanguageEnum{
		ZH_CN: ModifySecurityGroupRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ModifySecurityGroupRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ModifySecurityGroupRequestXLanguage) Value() string {
	return c.value
}

func (c ModifySecurityGroupRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModifySecurityGroupRequestXLanguage) UnmarshalJSON(b []byte) error {
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
