package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPolicyProtectRulesResponse Response Object
type ListPolicyProtectRulesResponse struct {

	// 策略总数
	Total *int32 `json:"total,omitempty"`

	// **参数解释：** 单个策略下特定类型防护规则列表 **约束限制：** 具体类型取决于参数rule_type **取值范围：** 具体返回值参考接口： - cc CC防护  ListCcPolicyRules - custom 精准防护 ListCustomPolicyRules - whiteblackip 黑白名单 ListWhiteblackipPolicyRules - geoip 地理位置防护 ListGeoIpPolicyRules - ip-reputation 威胁情报 ListIpReputationPolicyRules - antitamper 防篡改 ListAntiTamperPolicyRules - antileakage 防敏感信息泄露 ListAntileakagePolicyRules - ignore 全局白名单(原误报屏蔽) ListIgnorePolicyRules - privacy 隐私屏蔽 ListPrivacyPolicyRules **默认取值：** 不涉及
	Items          *[]interface{} `json:"items,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListPolicyProtectRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPolicyProtectRulesResponse struct{}"
	}

	return strings.Join([]string{"ListPolicyProtectRulesResponse", string(data)}, " ")
}
