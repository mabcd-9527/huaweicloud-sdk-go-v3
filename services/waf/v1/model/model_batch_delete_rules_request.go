package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchDeleteRulesRequest Request Object
type BatchDeleteRulesRequest struct {

	// **参数解释：** 需要删除的规则类型,目前支持cc,custom,whiteblackip,privacy,ignore,geoip,antitamper,antileakage,ip-reputation,llm-guards **约束限制：** 需要购买“大模型防火墙”服务后才可使用llm-guards **取值范围：** - cc CC防护 - custom 精准防护 - whiteblackip 黑白名单 - geoip 地理位置防护 - ip-reputation 威胁情报 - antitamper 防篡改 - antileakage 防敏感信息泄露 - ignore 全局白名单(原误报屏蔽) - privacy 隐私屏蔽 - llm-guards 大模型防火墙 **默认取值：** 不涉及
	RuleType BatchDeleteRulesRequestRuleType `json:"rule_type"`

	Body *PolicyRuleIdRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteRulesRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteRulesRequest", string(data)}, " ")
}

type BatchDeleteRulesRequestRuleType struct {
	value string
}

type BatchDeleteRulesRequestRuleTypeEnum struct {
	CC            BatchDeleteRulesRequestRuleType
	CUSTOM        BatchDeleteRulesRequestRuleType
	WHITEBLACKIP  BatchDeleteRulesRequestRuleType
	PRIVACY       BatchDeleteRulesRequestRuleType
	IGNORE        BatchDeleteRulesRequestRuleType
	GEOIP         BatchDeleteRulesRequestRuleType
	ANTITAMPER    BatchDeleteRulesRequestRuleType
	ANTILEAKAGE   BatchDeleteRulesRequestRuleType
	IP_REPUTATION BatchDeleteRulesRequestRuleType
	LLM_GUARDS    BatchDeleteRulesRequestRuleType
}

func GetBatchDeleteRulesRequestRuleTypeEnum() BatchDeleteRulesRequestRuleTypeEnum {
	return BatchDeleteRulesRequestRuleTypeEnum{
		CC: BatchDeleteRulesRequestRuleType{
			value: "cc",
		},
		CUSTOM: BatchDeleteRulesRequestRuleType{
			value: "custom",
		},
		WHITEBLACKIP: BatchDeleteRulesRequestRuleType{
			value: "whiteblackip",
		},
		PRIVACY: BatchDeleteRulesRequestRuleType{
			value: "privacy",
		},
		IGNORE: BatchDeleteRulesRequestRuleType{
			value: "ignore",
		},
		GEOIP: BatchDeleteRulesRequestRuleType{
			value: "geoip",
		},
		ANTITAMPER: BatchDeleteRulesRequestRuleType{
			value: "antitamper",
		},
		ANTILEAKAGE: BatchDeleteRulesRequestRuleType{
			value: "antileakage",
		},
		IP_REPUTATION: BatchDeleteRulesRequestRuleType{
			value: "ip-reputation",
		},
		LLM_GUARDS: BatchDeleteRulesRequestRuleType{
			value: "llm-guards",
		},
	}
}

func (c BatchDeleteRulesRequestRuleType) Value() string {
	return c.value
}

func (c BatchDeleteRulesRequestRuleType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchDeleteRulesRequestRuleType) UnmarshalJSON(b []byte) error {
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
