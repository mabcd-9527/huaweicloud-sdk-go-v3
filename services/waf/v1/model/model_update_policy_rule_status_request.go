package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdatePolicyRuleStatusRequest Request Object
type UpdatePolicyRuleStatusRequest struct {

	// **参数解释：** 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目ID。若需要查询当前用户所有企业项目绑定的资源信息，请传参all_granted_eps。 **约束限制：** 不涉及 **取值范围：**  - 0：代表default企业项目  - all_granted_eps：代表所有企业项目  - 其它企业项目ID：长度为36个字符 **默认取值：** 0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释：** 防护策略id，您可以通过调用查询防护策略列表（ListPolicy）获取策略id **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	PolicyId string `json:"policy_id"`

	// **参数解释：** 规则类型 **约束限制：** 不涉及 **取值范围：** - cc CC防护 - custom 精准防护 - whiteblackip 黑白名单 - geoip 地理位置防护 - ip-reputation 威胁情报 - antitamper 防篡改 - antileakage 防敏感信息泄露 - ignore 全局白名单(原误报屏蔽) - privacy 隐私屏蔽 **默认取值：** 不涉及
	RuleType UpdatePolicyRuleStatusRequestRuleType `json:"rule_type"`

	// **参数解释：** 规则id，通过对应规则类型的查询防护规则列表接口获取 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	RuleId string `json:"rule_id"`

	Body *UpdatePolicyRuleStatusRequestBody `json:"body,omitempty"`
}

func (o UpdatePolicyRuleStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePolicyRuleStatusRequest struct{}"
	}

	return strings.Join([]string{"UpdatePolicyRuleStatusRequest", string(data)}, " ")
}

type UpdatePolicyRuleStatusRequestRuleType struct {
	value string
}

type UpdatePolicyRuleStatusRequestRuleTypeEnum struct {
	CC            UpdatePolicyRuleStatusRequestRuleType
	CUSTOM        UpdatePolicyRuleStatusRequestRuleType
	WHITEBLACKIP  UpdatePolicyRuleStatusRequestRuleType
	PRIVACY       UpdatePolicyRuleStatusRequestRuleType
	IGNORE        UpdatePolicyRuleStatusRequestRuleType
	GEOIP         UpdatePolicyRuleStatusRequestRuleType
	ANTITAMPER    UpdatePolicyRuleStatusRequestRuleType
	ANTILEAKAGE   UpdatePolicyRuleStatusRequestRuleType
	IP_REPUTATION UpdatePolicyRuleStatusRequestRuleType
}

func GetUpdatePolicyRuleStatusRequestRuleTypeEnum() UpdatePolicyRuleStatusRequestRuleTypeEnum {
	return UpdatePolicyRuleStatusRequestRuleTypeEnum{
		CC: UpdatePolicyRuleStatusRequestRuleType{
			value: "cc",
		},
		CUSTOM: UpdatePolicyRuleStatusRequestRuleType{
			value: "custom",
		},
		WHITEBLACKIP: UpdatePolicyRuleStatusRequestRuleType{
			value: "whiteblackip",
		},
		PRIVACY: UpdatePolicyRuleStatusRequestRuleType{
			value: "privacy",
		},
		IGNORE: UpdatePolicyRuleStatusRequestRuleType{
			value: "ignore",
		},
		GEOIP: UpdatePolicyRuleStatusRequestRuleType{
			value: "geoip",
		},
		ANTITAMPER: UpdatePolicyRuleStatusRequestRuleType{
			value: "antitamper",
		},
		ANTILEAKAGE: UpdatePolicyRuleStatusRequestRuleType{
			value: "antileakage",
		},
		IP_REPUTATION: UpdatePolicyRuleStatusRequestRuleType{
			value: "ip-reputation",
		},
	}
}

func (c UpdatePolicyRuleStatusRequestRuleType) Value() string {
	return c.value
}

func (c UpdatePolicyRuleStatusRequestRuleType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdatePolicyRuleStatusRequestRuleType) UnmarshalJSON(b []byte) error {
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
