package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListPolicyProtectRulesRequest Request Object
type ListPolicyProtectRulesRequest struct {

	// **参数解释：** 策略id列表。策略id从\"查询防护策略列表\"(ListPolicy)接口获取，多个策略之间用“,”隔开 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Policyids *string `json:"policyids,omitempty"`

	// **参数解释：** 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目ID。若需要查询当前用户所有企业项目绑定的资源信息，请传参all_granted_eps。 **约束限制：** 不涉及 **取值范围：**  - 0：代表default企业项目  - all_granted_eps：代表所有企业项目  - 其它企业项目ID：长度为36个字符 **默认取值：** 0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释：** 分页查询时，返回第几页数据 **约束限制：** 不涉及 **取值范围：** page参数的实际有效范围取决于总数据量和pagesize的取值，不能大于总页数 **默认取值：** 1
	Page *int32 `json:"page,omitempty"`

	// **参数解释：** 分页查询时，每页包含的结果条数 **约束限制：** 不涉及 **取值范围：** [0, 总数据量] **默认取值：** 1000
	Pagesize *int32 `json:"pagesize,omitempty"`

	// **参数解释：** 需要查询的规则类型 **约束限制：** 不涉及 **取值范围：** - cc CC防护 - custom 精准防护 - whiteblackip 黑白名单 - geoip 地理位置防护 - ip-reputation 威胁情报 - antitamper 防篡改 - antileakage 防敏感信息泄露 - ignore 全局白名单(原误报屏蔽) - privacy 隐私屏蔽 **默认取值：** 不涉及
	RuleType ListPolicyProtectRulesRequestRuleType `json:"rule_type"`
}

func (o ListPolicyProtectRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPolicyProtectRulesRequest struct{}"
	}

	return strings.Join([]string{"ListPolicyProtectRulesRequest", string(data)}, " ")
}

type ListPolicyProtectRulesRequestRuleType struct {
	value string
}

type ListPolicyProtectRulesRequestRuleTypeEnum struct {
	CC            ListPolicyProtectRulesRequestRuleType
	CUSTOM        ListPolicyProtectRulesRequestRuleType
	WHITEBLACKIP  ListPolicyProtectRulesRequestRuleType
	PRIVACY       ListPolicyProtectRulesRequestRuleType
	IGNORE        ListPolicyProtectRulesRequestRuleType
	GEOIP         ListPolicyProtectRulesRequestRuleType
	ANTITAMPER    ListPolicyProtectRulesRequestRuleType
	ANTILEAKAGE   ListPolicyProtectRulesRequestRuleType
	IP_REPUTATION ListPolicyProtectRulesRequestRuleType
}

func GetListPolicyProtectRulesRequestRuleTypeEnum() ListPolicyProtectRulesRequestRuleTypeEnum {
	return ListPolicyProtectRulesRequestRuleTypeEnum{
		CC: ListPolicyProtectRulesRequestRuleType{
			value: "cc",
		},
		CUSTOM: ListPolicyProtectRulesRequestRuleType{
			value: "custom",
		},
		WHITEBLACKIP: ListPolicyProtectRulesRequestRuleType{
			value: "whiteblackip",
		},
		PRIVACY: ListPolicyProtectRulesRequestRuleType{
			value: "privacy",
		},
		IGNORE: ListPolicyProtectRulesRequestRuleType{
			value: "ignore",
		},
		GEOIP: ListPolicyProtectRulesRequestRuleType{
			value: "geoip",
		},
		ANTITAMPER: ListPolicyProtectRulesRequestRuleType{
			value: "antitamper",
		},
		ANTILEAKAGE: ListPolicyProtectRulesRequestRuleType{
			value: "antileakage",
		},
		IP_REPUTATION: ListPolicyProtectRulesRequestRuleType{
			value: "ip-reputation",
		},
	}
}

func (c ListPolicyProtectRulesRequestRuleType) Value() string {
	return c.value
}

func (c ListPolicyProtectRulesRequestRuleType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPolicyProtectRulesRequestRuleType) UnmarshalJSON(b []byte) error {
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
