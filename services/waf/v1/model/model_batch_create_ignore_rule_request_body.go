package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateIgnoreRuleRequestBody 创建误报屏蔽规则请求体
type BatchCreateIgnoreRuleRequestBody struct {

	// 防护域名或防护网站，数组长度为0时，代表规则对全部域名或防护网站生效。当防护域名的接入模式为云模式-ELB接入时，该参数需以<域名>:\\<id\\>格式填写（如www.example.com:b061fb5b-8ea0-4357-b0f4-cb6178ab378a），若域名绑定的负载均衡器（ELB）下所有监听器都接入WAF防护，填入的id为负载均衡器（ELB）id，否则填入的id为指定监听器id；可通过ShowPremiumHost接口查询与该域名绑定的ELB实例id，在帮助中心-ELB服务-API文档下查询监听器id
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

	// **参数解释：** 添加规则的策略id列表。策略id从\"查询防护策略列表\"(ListPolicy)接口获取，多个策略之间用“,”隔开 **约束限制：** 不能为空 **取值范围：** 不涉及 **默认取值：** 不涉及
	PolicyIds []string `json:"policy_ids"`
}

func (o BatchCreateIgnoreRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateIgnoreRuleRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreateIgnoreRuleRequestBody", string(data)}, " ")
}
