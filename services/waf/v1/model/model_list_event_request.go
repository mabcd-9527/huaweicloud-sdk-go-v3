package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListEventRequest Request Object
type ListEventRequest struct {

	// **参数解释：** 客户端IP所属地理位置展示语言，默认值为en-us **约束限制：** 不涉及 **取值范围：** - zh-cn 中文 - en-us 英文 **默认取值：** en-us
	XLanguage *ListEventRequestXLanguage `json:"X-Language,omitempty"`

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id。若需要查询当前用户所有企业项目绑定的资源信息，请传参all_granted_eps。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释：** 查询日志的时间范围，recent参数与from、to必须使用其中一个。当同时使用recent参数与from、to时，以recent参数为准 **约束限制：** 不涉及 **取值范围：**  - yesterday：昨天  - today：今天  - 3days：近3天   - 1week：近7天   - 1month：近30天  **默认取值：** 不涉及
	Recent *ListEventRequestRecent `json:"recent,omitempty"`

	// **参数解释：** 起始时间(毫秒时间戳)，需要和to同时使用 **约束限制：** from <= to **取值范围：** from ~ to 最大范围30天 **默认取值：** 不涉及
	From *int64 `json:"from,omitempty"`

	// **参数解释：** 结束时间(毫秒时间戳)，需要和from同时使用 **约束限制：** from ~ to 最大范围30天 **取值范围：** 不能超过当天的结束时间 **默认取值：** 不涉及
	To *int64 `json:"to,omitempty"`

	// **参数解释：** 防护事件id列表，支持模糊查询 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Ids *[]string `json:"ids,omitempty"`

	// **参数解释：** 防护事件id列表（排除搜索），支持模糊查询 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Nids *[]string `json:"nids,omitempty"`

	// **参数解释：** 攻击类型 **约束限制：** 不涉及 **取值范围：** - xss：XSS攻击  - botm：BOT攻击 - webshell：网站木马  - vuln：其他漏洞攻击 - sqli：sql注入攻击  - robot：恶意爬虫  - rfi：远程文件包含  - rce：远程代码执行 - ptr：目录遍历 - lfi：本地文件包含 - antileakage：网站信息泄漏  - iprank：IP信誉库 - custom_whiteblackip：IP黑白名单 - custom_whiteip：白名单 - custom_blackip：黑名单 - custom_robot：扫描器爬虫 - custom_geoip：地理访问控制 - custom_idc_ip：IDC情报 - custom_custom：精准防护  - cmdi：命令注入攻击  - cc：cc攻击  - antitamper：网页防篡改  - anticrawler：网站反爬虫   - third_bot_river：第三方反爬虫 - antiscan_high_freq_scan：高频扫描封禁 - antiscan_dir_traversal：目录遍历防护 - illegal：非法请求 - followed_action：攻击惩罚 - advanced_bot：BOT管理 - llm_prompt_injection：提示词注入攻击 - llm_prompt_sensitive：提示词违规 - llm_response_sensitive：响应违规 **默认取值：** 不涉及
	Attacks *[]ListEventRequestAttacks `json:"attacks,omitempty"`

	// **参数解释：** 攻击类型（排除搜索） **约束限制：** 不涉及 **取值范围：** - xss：XSS攻击  - botm：BOT攻击 - webshell：网站木马  - vuln：其他漏洞攻击 - sqli：sql注入攻击  - robot：恶意爬虫  - rfi：远程文件包含  - rce：远程代码执行 - ptr：目录遍历 - lfi：本地文件包含 - antileakage：网站信息泄漏  - iprank：IP信誉库 - custom_whiteblackip：IP黑白名单 - custom_whiteip：白名单 - custom_blackip：黑名单 - custom_robot：扫描器爬虫 - custom_geoip：地理访问控制 - custom_idc_ip：IDC情报 - custom_custom：精准防护  - cmdi：命令注入攻击  - cc：cc攻击  - antitamper：网页防篡改  - anticrawler：网站反爬虫   - third_bot_river：第三方反爬虫 - antiscan_high_freq_scan：高频扫描封禁 - antiscan_dir_traversal：目录遍历防护 - illegal：非法请求 - followed_action：攻击惩罚 - advanced_bot：BOT管理 - llm_prompt_injection：提示词注入攻击 - llm_prompt_sensitive：提示词违规 - llm_response_sensitive：响应违规 **默认取值：** 不涉及
	Nattacks *[]ListEventRequestNattacks `json:"nattacks,omitempty"`

	// **参数解释：** 规则id列表 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Rules *[]string `json:"rules,omitempty"`

	// **参数解释：** 规则id列表（排除搜索） **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Nrules *[]string `json:"nrules,omitempty"`

	// **参数解释：** 客户端IP列表 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Sips *[]string `json:"sips,omitempty"`

	// **参数解释：** 客户端IP列表（排除搜索） **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Nsips *[]string `json:"nsips,omitempty"`

	// **参数解释：** 客户端IP，当query_mode为\"equal\"时为精确查询，否则模糊查询 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Sip *string `json:"sip,omitempty"`

	// **参数解释：** url列表 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Urls *[]string `json:"urls,omitempty"`

	// **参数解释：** url列表（排除搜索） **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Nurls *[]string `json:"nurls,omitempty"`

	// **参数解释：** URL，当query_mode为\"equal\"时为精确查询，否则模糊查询 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Url *string `json:"url,omitempty"`

	// **参数解释：** 防护动作列表 **约束限制：** 不涉及 **取值范围：** - block：拦截 - pass：放行 - log：仅记录 - captcha：人机验证 - cache：不匹配 - mask：过滤 - js_challenge：JS挑战 - advanced_captcha：高级人机验证 - abort_response：中断响应 - desensitize：脱敏 **默认取值：** 不涉及
	Actions *[]ListEventRequestActions `json:"actions,omitempty"`

	// **参数解释：** 防护动作列表（排除搜索） **约束限制：** 不涉及 **取值范围：** - block：拦截 - pass：放行 - log：仅记录 - captcha：人机验证 - cache：不匹配 - mask：过滤 - js_challenge：JS挑战 - advanced_captcha：高级人机验证 - abort_response：中断响应 - desensitize：脱敏 **默认取值：** 不涉及
	Nactions *[]ListEventRequestNactions `json:"nactions,omitempty"`

	// **参数解释：** 域名，支持模糊查询 **约束限制：** domain和ndomain不可同时查询，当两个都存在时以domain为准 **取值范围：** 不涉及 **默认取值：** 不涉及
	Domain *string `json:"domain,omitempty"`

	// **参数解释：** 域名（排除搜索），支持模糊查询 **约束限制：** domain和ndomain不可同时查询，当两个都存在时以domain为准 **取值范围：** 不涉及 **默认取值：** 不涉及
	Ndomain *string `json:"ndomain,omitempty"`

	// **参数解释：** 域名列表 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Domains *[]string `json:"domains,omitempty"`

	// **参数解释：** 客户端IP所属国家列表 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	IpCountries *[]string `json:"ip_countries,omitempty"`

	// **参数解释：** 客户端IP所属国家列表（排除搜索） **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	NipCountries *[]string `json:"nip_countries,omitempty"`

	// **参数解释：** 客户端IP所属省份列表，仅中国省份生效 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	IpRegions *[]string `json:"ip_regions,omitempty"`

	// **参数解释：** 客户端IP所属身份列表（排除搜索），仅中国省份生效 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	NipRegions *[]string `json:"nip_regions,omitempty"`

	// **参数解释：** 响应码列表 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ResponseCodes *[]string `json:"response_codes,omitempty"`

	// **参数解释：** 恶意负载（被WAF识别的攻击片段）： Web 基础防护(SQL注入、XSS、命令注入等)：被WAF识别的攻击片段 CC 攻击：命中规则的请求次数 精准防护、IP黑白名单、地理访问控制：空 攻击惩罚：命中攻击惩罚的用户标识 恶意爬虫：命中规则的 User-Agent 字段 网页反爬虫：JS 脚本事件：js_verified（JS 脚本验证通过事件）和 js_challenge（发送 JS 验证内容事件）。如果请求验证失败则为空。 网站信息泄露：敏感信息过滤为过滤类型，既电话号码,电子邮箱,身份证号；响应码拦截则为拦截的响应码值。 BOT攻击：命中规则的User-Agent等异常请求特征，或AI行为检测结果的评分细节 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Payload *string `json:"payload,omitempty"`

	// **参数解释：** 域名id列表，从获取防护网站列表（ListHost）接口获取域名id **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Hosts *[]string `json:"hosts,omitempty"`

	// **参数解释：** 引擎实例id列表 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Instances *[]string `json:"instances,omitempty"`

	// **参数解释：** 分页查询时，返回第几页数据 **约束限制：** 不涉及 **取值范围：** page参数的实际有效范围取决于总数据量和pagesize的取值，不能大于总页数 **默认取值：** 1
	Page *int32 `json:"page,omitempty"`

	// **参数解释：** 分页查询时，每页包含的结果条数 **约束限制：** 不涉及 **取值范围：** [0, 总数据量] **默认取值：** 10
	Pagesize *int32 `json:"pagesize,omitempty"`

	// **参数解释：** 排序字段，默认attack_time，选择其他字段时，会按照指定字段和attack_time共同排序 **约束限制：** 不涉及 **取值范围：** - attack_time 攻击时间 - sort_ip 客户端IP - host 域名 - geo_str 地理位置 - component 应用组件 - rule 规则ID - attack 事件类型（攻击类型） **默认取值：** attack_time
	SortKey *ListEventRequestSortKey `json:"sort_key,omitempty"`

	// **参数解释：** 排序方向 **约束限制：** 不涉及 **取值范围：** - desc 降序 - asc 升序 **默认取值：** desc
	SortDirection *ListEventRequestSortDirection `json:"sort_direction,omitempty"`

	// **参数解释：** 查询模式，仅影响参数sip、url **约束限制：** 不涉及 **取值范围：** - equal 精确查询 - include 模糊查询 **默认取值：** include
	QueryMode *ListEventRequestQueryMode `json:"query_mode,omitempty"`
}

func (o ListEventRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEventRequest struct{}"
	}

	return strings.Join([]string{"ListEventRequest", string(data)}, " ")
}

type ListEventRequestXLanguage struct {
	value string
}

type ListEventRequestXLanguageEnum struct {
	ZH_CN ListEventRequestXLanguage
	EN_US ListEventRequestXLanguage
}

func GetListEventRequestXLanguageEnum() ListEventRequestXLanguageEnum {
	return ListEventRequestXLanguageEnum{
		ZH_CN: ListEventRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListEventRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListEventRequestXLanguage) Value() string {
	return c.value
}

func (c ListEventRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEventRequestXLanguage) UnmarshalJSON(b []byte) error {
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

type ListEventRequestRecent struct {
	value string
}

type ListEventRequestRecentEnum struct {
	YESTERDAY ListEventRequestRecent
	TODAY     ListEventRequestRecent
	E_3DAYS   ListEventRequestRecent
	E_1WEEK   ListEventRequestRecent
	E_1MONTH  ListEventRequestRecent
}

func GetListEventRequestRecentEnum() ListEventRequestRecentEnum {
	return ListEventRequestRecentEnum{
		YESTERDAY: ListEventRequestRecent{
			value: "yesterday",
		},
		TODAY: ListEventRequestRecent{
			value: "today",
		},
		E_3DAYS: ListEventRequestRecent{
			value: "3days",
		},
		E_1WEEK: ListEventRequestRecent{
			value: "1week",
		},
		E_1MONTH: ListEventRequestRecent{
			value: "1month",
		},
	}
}

func (c ListEventRequestRecent) Value() string {
	return c.value
}

func (c ListEventRequestRecent) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEventRequestRecent) UnmarshalJSON(b []byte) error {
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

type ListEventRequestAttacks struct {
	value string
}

type ListEventRequestAttacksEnum struct {
	XSS                     ListEventRequestAttacks
	BOTM                    ListEventRequestAttacks
	WEBSHELL                ListEventRequestAttacks
	VULN                    ListEventRequestAttacks
	SQLI                    ListEventRequestAttacks
	ROBOT                   ListEventRequestAttacks
	RFI                     ListEventRequestAttacks
	RCE                     ListEventRequestAttacks
	PTR                     ListEventRequestAttacks
	LFI                     ListEventRequestAttacks
	ANTILEAKAGE             ListEventRequestAttacks
	IPRANK                  ListEventRequestAttacks
	CUSTOM_WHITEBLACKIP     ListEventRequestAttacks
	CUSTOM_WHITEIP          ListEventRequestAttacks
	CUSTOM_BLACKIP          ListEventRequestAttacks
	CUSTOM_ROBOT            ListEventRequestAttacks
	CUSTOM_GEOIP            ListEventRequestAttacks
	CUSTOM_IDC_IP           ListEventRequestAttacks
	CUSTOM_CUSTOM           ListEventRequestAttacks
	CMDI                    ListEventRequestAttacks
	CC                      ListEventRequestAttacks
	ANTITAMPER              ListEventRequestAttacks
	ANTICRAWLER             ListEventRequestAttacks
	THIRD_BOT_RIVER         ListEventRequestAttacks
	ANTISCAN_HIGH_FREQ_SCAN ListEventRequestAttacks
	ANTISCAN_DIR_TRAVERSAL  ListEventRequestAttacks
	ILLEGAL                 ListEventRequestAttacks
	FOLLOWED_ACTION         ListEventRequestAttacks
	ADVANCED_BOT            ListEventRequestAttacks
	LLM_PROMPT_INJECTION    ListEventRequestAttacks
	LLM_PROMPT_SENSITIVE    ListEventRequestAttacks
	LLM_RESPONSE_SENSITIVE  ListEventRequestAttacks
}

func GetListEventRequestAttacksEnum() ListEventRequestAttacksEnum {
	return ListEventRequestAttacksEnum{
		XSS: ListEventRequestAttacks{
			value: "xss",
		},
		BOTM: ListEventRequestAttacks{
			value: "botm",
		},
		WEBSHELL: ListEventRequestAttacks{
			value: "webshell",
		},
		VULN: ListEventRequestAttacks{
			value: "vuln",
		},
		SQLI: ListEventRequestAttacks{
			value: "sqli",
		},
		ROBOT: ListEventRequestAttacks{
			value: "robot",
		},
		RFI: ListEventRequestAttacks{
			value: "rfi",
		},
		RCE: ListEventRequestAttacks{
			value: "rce",
		},
		PTR: ListEventRequestAttacks{
			value: "ptr",
		},
		LFI: ListEventRequestAttacks{
			value: "lfi",
		},
		ANTILEAKAGE: ListEventRequestAttacks{
			value: "antileakage",
		},
		IPRANK: ListEventRequestAttacks{
			value: "iprank",
		},
		CUSTOM_WHITEBLACKIP: ListEventRequestAttacks{
			value: "custom_whiteblackip",
		},
		CUSTOM_WHITEIP: ListEventRequestAttacks{
			value: "custom_whiteip",
		},
		CUSTOM_BLACKIP: ListEventRequestAttacks{
			value: "custom_blackip",
		},
		CUSTOM_ROBOT: ListEventRequestAttacks{
			value: "custom_robot",
		},
		CUSTOM_GEOIP: ListEventRequestAttacks{
			value: "custom_geoip",
		},
		CUSTOM_IDC_IP: ListEventRequestAttacks{
			value: "custom_idc_ip",
		},
		CUSTOM_CUSTOM: ListEventRequestAttacks{
			value: "custom_custom",
		},
		CMDI: ListEventRequestAttacks{
			value: "cmdi",
		},
		CC: ListEventRequestAttacks{
			value: "cc",
		},
		ANTITAMPER: ListEventRequestAttacks{
			value: "antitamper",
		},
		ANTICRAWLER: ListEventRequestAttacks{
			value: "anticrawler",
		},
		THIRD_BOT_RIVER: ListEventRequestAttacks{
			value: "third_bot_river",
		},
		ANTISCAN_HIGH_FREQ_SCAN: ListEventRequestAttacks{
			value: "antiscan_high_freq_scan",
		},
		ANTISCAN_DIR_TRAVERSAL: ListEventRequestAttacks{
			value: "antiscan_dir_traversal",
		},
		ILLEGAL: ListEventRequestAttacks{
			value: "illegal",
		},
		FOLLOWED_ACTION: ListEventRequestAttacks{
			value: "followed_action",
		},
		ADVANCED_BOT: ListEventRequestAttacks{
			value: "advanced_bot",
		},
		LLM_PROMPT_INJECTION: ListEventRequestAttacks{
			value: "llm_prompt_injection",
		},
		LLM_PROMPT_SENSITIVE: ListEventRequestAttacks{
			value: "llm_prompt_sensitive",
		},
		LLM_RESPONSE_SENSITIVE: ListEventRequestAttacks{
			value: "llm_response_sensitive",
		},
	}
}

func (c ListEventRequestAttacks) Value() string {
	return c.value
}

func (c ListEventRequestAttacks) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEventRequestAttacks) UnmarshalJSON(b []byte) error {
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

type ListEventRequestNattacks struct {
	value string
}

type ListEventRequestNattacksEnum struct {
	XSS                     ListEventRequestNattacks
	BOTM                    ListEventRequestNattacks
	WEBSHELL                ListEventRequestNattacks
	VULN                    ListEventRequestNattacks
	SQLI                    ListEventRequestNattacks
	ROBOT                   ListEventRequestNattacks
	RFI                     ListEventRequestNattacks
	RCE                     ListEventRequestNattacks
	PTR                     ListEventRequestNattacks
	LFI                     ListEventRequestNattacks
	ANTILEAKAGE             ListEventRequestNattacks
	IPRANK                  ListEventRequestNattacks
	CUSTOM_WHITEBLACKIP     ListEventRequestNattacks
	CUSTOM_WHITEIP          ListEventRequestNattacks
	CUSTOM_BLACKIP          ListEventRequestNattacks
	CUSTOM_ROBOT            ListEventRequestNattacks
	CUSTOM_GEOIP            ListEventRequestNattacks
	CUSTOM_IDC_IP           ListEventRequestNattacks
	CUSTOM_CUSTOM           ListEventRequestNattacks
	CMDI                    ListEventRequestNattacks
	CC                      ListEventRequestNattacks
	ANTITAMPER              ListEventRequestNattacks
	ANTICRAWLER             ListEventRequestNattacks
	THIRD_BOT_RIVER         ListEventRequestNattacks
	ANTISCAN_HIGH_FREQ_SCAN ListEventRequestNattacks
	ANTISCAN_DIR_TRAVERSAL  ListEventRequestNattacks
	ILLEGAL                 ListEventRequestNattacks
	FOLLOWED_ACTION         ListEventRequestNattacks
	ADVANCED_BOT            ListEventRequestNattacks
	LLM_PROMPT_INJECTION    ListEventRequestNattacks
	LLM_PROMPT_SENSITIVE    ListEventRequestNattacks
	LLM_RESPONSE_SENSITIVE  ListEventRequestNattacks
}

func GetListEventRequestNattacksEnum() ListEventRequestNattacksEnum {
	return ListEventRequestNattacksEnum{
		XSS: ListEventRequestNattacks{
			value: "xss",
		},
		BOTM: ListEventRequestNattacks{
			value: "botm",
		},
		WEBSHELL: ListEventRequestNattacks{
			value: "webshell",
		},
		VULN: ListEventRequestNattacks{
			value: "vuln",
		},
		SQLI: ListEventRequestNattacks{
			value: "sqli",
		},
		ROBOT: ListEventRequestNattacks{
			value: "robot",
		},
		RFI: ListEventRequestNattacks{
			value: "rfi",
		},
		RCE: ListEventRequestNattacks{
			value: "rce",
		},
		PTR: ListEventRequestNattacks{
			value: "ptr",
		},
		LFI: ListEventRequestNattacks{
			value: "lfi",
		},
		ANTILEAKAGE: ListEventRequestNattacks{
			value: "antileakage",
		},
		IPRANK: ListEventRequestNattacks{
			value: "iprank",
		},
		CUSTOM_WHITEBLACKIP: ListEventRequestNattacks{
			value: "custom_whiteblackip",
		},
		CUSTOM_WHITEIP: ListEventRequestNattacks{
			value: "custom_whiteip",
		},
		CUSTOM_BLACKIP: ListEventRequestNattacks{
			value: "custom_blackip",
		},
		CUSTOM_ROBOT: ListEventRequestNattacks{
			value: "custom_robot",
		},
		CUSTOM_GEOIP: ListEventRequestNattacks{
			value: "custom_geoip",
		},
		CUSTOM_IDC_IP: ListEventRequestNattacks{
			value: "custom_idc_ip",
		},
		CUSTOM_CUSTOM: ListEventRequestNattacks{
			value: "custom_custom",
		},
		CMDI: ListEventRequestNattacks{
			value: "cmdi",
		},
		CC: ListEventRequestNattacks{
			value: "cc",
		},
		ANTITAMPER: ListEventRequestNattacks{
			value: "antitamper",
		},
		ANTICRAWLER: ListEventRequestNattacks{
			value: "anticrawler",
		},
		THIRD_BOT_RIVER: ListEventRequestNattacks{
			value: "third_bot_river",
		},
		ANTISCAN_HIGH_FREQ_SCAN: ListEventRequestNattacks{
			value: "antiscan_high_freq_scan",
		},
		ANTISCAN_DIR_TRAVERSAL: ListEventRequestNattacks{
			value: "antiscan_dir_traversal",
		},
		ILLEGAL: ListEventRequestNattacks{
			value: "illegal",
		},
		FOLLOWED_ACTION: ListEventRequestNattacks{
			value: "followed_action",
		},
		ADVANCED_BOT: ListEventRequestNattacks{
			value: "advanced_bot",
		},
		LLM_PROMPT_INJECTION: ListEventRequestNattacks{
			value: "llm_prompt_injection",
		},
		LLM_PROMPT_SENSITIVE: ListEventRequestNattacks{
			value: "llm_prompt_sensitive",
		},
		LLM_RESPONSE_SENSITIVE: ListEventRequestNattacks{
			value: "llm_response_sensitive",
		},
	}
}

func (c ListEventRequestNattacks) Value() string {
	return c.value
}

func (c ListEventRequestNattacks) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEventRequestNattacks) UnmarshalJSON(b []byte) error {
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

type ListEventRequestActions struct {
	value string
}

type ListEventRequestActionsEnum struct {
	BLOCK            ListEventRequestActions
	PASS             ListEventRequestActions
	LOG              ListEventRequestActions
	CAPTCHA          ListEventRequestActions
	CACHE            ListEventRequestActions
	MASK             ListEventRequestActions
	JS_CHALLENGE     ListEventRequestActions
	ADVANCED_CAPTCHA ListEventRequestActions
	ABORT_RESPONSE   ListEventRequestActions
	DESENSITIZE      ListEventRequestActions
}

func GetListEventRequestActionsEnum() ListEventRequestActionsEnum {
	return ListEventRequestActionsEnum{
		BLOCK: ListEventRequestActions{
			value: "block",
		},
		PASS: ListEventRequestActions{
			value: "pass",
		},
		LOG: ListEventRequestActions{
			value: "log",
		},
		CAPTCHA: ListEventRequestActions{
			value: "captcha",
		},
		CACHE: ListEventRequestActions{
			value: "cache",
		},
		MASK: ListEventRequestActions{
			value: "mask",
		},
		JS_CHALLENGE: ListEventRequestActions{
			value: "js_challenge",
		},
		ADVANCED_CAPTCHA: ListEventRequestActions{
			value: "advanced_captcha",
		},
		ABORT_RESPONSE: ListEventRequestActions{
			value: "abort_response",
		},
		DESENSITIZE: ListEventRequestActions{
			value: "desensitize",
		},
	}
}

func (c ListEventRequestActions) Value() string {
	return c.value
}

func (c ListEventRequestActions) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEventRequestActions) UnmarshalJSON(b []byte) error {
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

type ListEventRequestNactions struct {
	value string
}

type ListEventRequestNactionsEnum struct {
	BLOCK            ListEventRequestNactions
	PASS             ListEventRequestNactions
	LOG              ListEventRequestNactions
	CAPTCHA          ListEventRequestNactions
	CACHE            ListEventRequestNactions
	MASK             ListEventRequestNactions
	JS_CHALLENGE     ListEventRequestNactions
	ADVANCED_CAPTCHA ListEventRequestNactions
	ABORT_RESPONSE   ListEventRequestNactions
	DESENSITIZE      ListEventRequestNactions
}

func GetListEventRequestNactionsEnum() ListEventRequestNactionsEnum {
	return ListEventRequestNactionsEnum{
		BLOCK: ListEventRequestNactions{
			value: "block",
		},
		PASS: ListEventRequestNactions{
			value: "pass",
		},
		LOG: ListEventRequestNactions{
			value: "log",
		},
		CAPTCHA: ListEventRequestNactions{
			value: "captcha",
		},
		CACHE: ListEventRequestNactions{
			value: "cache",
		},
		MASK: ListEventRequestNactions{
			value: "mask",
		},
		JS_CHALLENGE: ListEventRequestNactions{
			value: "js_challenge",
		},
		ADVANCED_CAPTCHA: ListEventRequestNactions{
			value: "advanced_captcha",
		},
		ABORT_RESPONSE: ListEventRequestNactions{
			value: "abort_response",
		},
		DESENSITIZE: ListEventRequestNactions{
			value: "desensitize",
		},
	}
}

func (c ListEventRequestNactions) Value() string {
	return c.value
}

func (c ListEventRequestNactions) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEventRequestNactions) UnmarshalJSON(b []byte) error {
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

type ListEventRequestSortKey struct {
	value string
}

type ListEventRequestSortKeyEnum struct {
	ATTACK_TIME ListEventRequestSortKey
	SORT_IP     ListEventRequestSortKey
	HOST        ListEventRequestSortKey
	GEO_STR     ListEventRequestSortKey
	COMPONENT   ListEventRequestSortKey
	RULE        ListEventRequestSortKey
	ATTACK      ListEventRequestSortKey
}

func GetListEventRequestSortKeyEnum() ListEventRequestSortKeyEnum {
	return ListEventRequestSortKeyEnum{
		ATTACK_TIME: ListEventRequestSortKey{
			value: "attack_time",
		},
		SORT_IP: ListEventRequestSortKey{
			value: "sort_ip",
		},
		HOST: ListEventRequestSortKey{
			value: "host",
		},
		GEO_STR: ListEventRequestSortKey{
			value: "geo_str",
		},
		COMPONENT: ListEventRequestSortKey{
			value: "component",
		},
		RULE: ListEventRequestSortKey{
			value: "rule",
		},
		ATTACK: ListEventRequestSortKey{
			value: "attack",
		},
	}
}

func (c ListEventRequestSortKey) Value() string {
	return c.value
}

func (c ListEventRequestSortKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEventRequestSortKey) UnmarshalJSON(b []byte) error {
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

type ListEventRequestSortDirection struct {
	value string
}

type ListEventRequestSortDirectionEnum struct {
	DESC ListEventRequestSortDirection
	ASC  ListEventRequestSortDirection
}

func GetListEventRequestSortDirectionEnum() ListEventRequestSortDirectionEnum {
	return ListEventRequestSortDirectionEnum{
		DESC: ListEventRequestSortDirection{
			value: "desc",
		},
		ASC: ListEventRequestSortDirection{
			value: "asc",
		},
	}
}

func (c ListEventRequestSortDirection) Value() string {
	return c.value
}

func (c ListEventRequestSortDirection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEventRequestSortDirection) UnmarshalJSON(b []byte) error {
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

type ListEventRequestQueryMode struct {
	value string
}

type ListEventRequestQueryModeEnum struct {
	EQUAL   ListEventRequestQueryMode
	INCLUDE ListEventRequestQueryMode
}

func GetListEventRequestQueryModeEnum() ListEventRequestQueryModeEnum {
	return ListEventRequestQueryModeEnum{
		EQUAL: ListEventRequestQueryMode{
			value: "equal",
		},
		INCLUDE: ListEventRequestQueryMode{
			value: "include",
		},
	}
}

func (c ListEventRequestQueryMode) Value() string {
	return c.value
}

func (c ListEventRequestQueryMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListEventRequestQueryMode) UnmarshalJSON(b []byte) error {
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
