package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ConfirmPolicyIpReputationMapRequest Request Object
type ConfirmPolicyIpReputationMapRequest struct {

	// **参数解释：** 语言类型 **约束限制：** 不涉及 **取值范围：** - cn 中文 - en 英文 **默认取值：** cn
	Lang *ConfirmPolicyIpReputationMapRequestLang `json:"lang,omitempty"`

	// **参数解释：** 防护选项的详细信息的类型，当前仅支持“idc”。 **约束限制：** 不涉及 **取值范围：** - idc **默认取值：** 不涉及
	Type ConfirmPolicyIpReputationMapRequestType `json:"type"`
}

func (o ConfirmPolicyIpReputationMapRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmPolicyIpReputationMapRequest struct{}"
	}

	return strings.Join([]string{"ConfirmPolicyIpReputationMapRequest", string(data)}, " ")
}

type ConfirmPolicyIpReputationMapRequestLang struct {
	value string
}

type ConfirmPolicyIpReputationMapRequestLangEnum struct {
	CN ConfirmPolicyIpReputationMapRequestLang
	EN ConfirmPolicyIpReputationMapRequestLang
}

func GetConfirmPolicyIpReputationMapRequestLangEnum() ConfirmPolicyIpReputationMapRequestLangEnum {
	return ConfirmPolicyIpReputationMapRequestLangEnum{
		CN: ConfirmPolicyIpReputationMapRequestLang{
			value: "cn",
		},
		EN: ConfirmPolicyIpReputationMapRequestLang{
			value: "en",
		},
	}
}

func (c ConfirmPolicyIpReputationMapRequestLang) Value() string {
	return c.value
}

func (c ConfirmPolicyIpReputationMapRequestLang) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConfirmPolicyIpReputationMapRequestLang) UnmarshalJSON(b []byte) error {
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

type ConfirmPolicyIpReputationMapRequestType struct {
	value string
}

type ConfirmPolicyIpReputationMapRequestTypeEnum struct {
	IDC ConfirmPolicyIpReputationMapRequestType
}

func GetConfirmPolicyIpReputationMapRequestTypeEnum() ConfirmPolicyIpReputationMapRequestTypeEnum {
	return ConfirmPolicyIpReputationMapRequestTypeEnum{
		IDC: ConfirmPolicyIpReputationMapRequestType{
			value: "idc",
		},
	}
}

func (c ConfirmPolicyIpReputationMapRequestType) Value() string {
	return c.value
}

func (c ConfirmPolicyIpReputationMapRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConfirmPolicyIpReputationMapRequestType) UnmarshalJSON(b []byte) error {
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
