package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateIgnoreRuleRequestBody 更新误报屏蔽规则请求体
type BatchUpdateIgnoreRuleRequestBody struct {

	// 防护域名或防护网站，数组长度为0时，代表规则对全部域名或防护网站生效
	Domain []string `json:"domain"`

	// 条件列表
	Conditions []CreateCondition `json:"conditions"`

	// 固定值为1,代表v2版本误报屏蔽规则，v1版本仅用于兼容旧版本，不支持创建
	Mode int32 `json:"mode"`

	// **参数解释：** 需要屏蔽的规则 **约束限制：** 参数值根据\"不检测模块\"变化 **取值范围：** 不检测模块: - 所有模块: bypass - Web基础防护模块按照规则类型划分:   - ID: 内置规则id，通过ListWebBasicProtectionRules接口获取ID，多个id以分号;分隔，比如：\"000000;111111\"   - 类别: 多个类型以分号;分隔，比如：\"xss;webshell\"     - xss：XSS攻击     - webshell：网站木马     - vuln：其他类型攻击     - sqli：SQL注入攻击     - robot：恶意爬虫     - rfi：远程文件包含     - lfi：本地文件包含     - cmdi：命令注入攻击   - 所有内置规则 - 非法请求: illegal **默认取值：** 不涉及
	Rule string `json:"rule"`

	Advanced *IgnoreAdvanced `json:"advanced,omitempty"`

	// 屏蔽规则描述
	Description *string `json:"description,omitempty"`

	// **参数解释：** 策略ID和规则id数组，规则ID与规则所属的策略ID **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	PolicyRuleIds []BatchUpdateIgnoreRuleRequestBodyPolicyRuleIds `json:"policy_rule_ids"`
}

func (o BatchUpdateIgnoreRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateIgnoreRuleRequestBody struct{}"
	}

	return strings.Join([]string{"BatchUpdateIgnoreRuleRequestBody", string(data)}, " ")
}
