package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowWebProtectionRuleRequest Request Object
type ShowWebProtectionRuleRequest struct {

	// **参数解释：** 语言类型 **约束限制：** 不涉及 **取值范围：** - zh-cn 中文 - en-us 英文 **默认取值：** zh-cn
	XLanguage *ShowWebProtectionRuleRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释：** 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目ID。若需要查询当前用户所有企业项目绑定的资源信息，请传参all_granted_eps。 **约束限制：** 不涉及 **取值范围：**  - 0：代表default企业项目  - all_granted_eps：代表所有企业项目  - 其它企业项目ID：长度为36个字符 **默认取值：** 0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释：** 基础防护规则id，通过查询内置规则集列表（ListWebBasicProtectionRules）接口获取 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	RuleId string `json:"rule_id"`
}

func (o ShowWebProtectionRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWebProtectionRuleRequest struct{}"
	}

	return strings.Join([]string{"ShowWebProtectionRuleRequest", string(data)}, " ")
}

type ShowWebProtectionRuleRequestXLanguage struct {
	value string
}

type ShowWebProtectionRuleRequestXLanguageEnum struct {
	ZH_CN ShowWebProtectionRuleRequestXLanguage
	EN_US ShowWebProtectionRuleRequestXLanguage
}

func GetShowWebProtectionRuleRequestXLanguageEnum() ShowWebProtectionRuleRequestXLanguageEnum {
	return ShowWebProtectionRuleRequestXLanguageEnum{
		ZH_CN: ShowWebProtectionRuleRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowWebProtectionRuleRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowWebProtectionRuleRequestXLanguage) Value() string {
	return c.value
}

func (c ShowWebProtectionRuleRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowWebProtectionRuleRequestXLanguage) UnmarshalJSON(b []byte) error {
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
