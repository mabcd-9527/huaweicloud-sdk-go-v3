package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateCondition 条件列表参数较为复杂，存在级联关系，建议同时使用控制台上的添加误报屏蔽规则，单击F12键查看路径后缀为/ignore，方法为POST的请求参数，便于理解参数的填写
type CreateCondition struct {

	// **参数解释：** 字段类型 **约束限制：** 不涉及 **取值范围：** - url: 路径 - custom_geoip: 客户端IP所属的地理位置 - robot: 已知特征爬虫 - user-agent: User Agent - ip: IPv4 - ipv6: IPv6 - params: Params - cookie: Cookie - referer: Referer - header: Header - method: Method - request_line: Request Line - request: Request - protocol: Protocol - request_body: Request Body  **默认取值：** 不涉及
	Category CreateConditionCategory `json:"category"`

	// **参数解释：** 子字段 **约束限制：** 随字段类型变化 **取值范围：** - custom_geoip:    - v4：ipv4   - v6：ipv6   - any：ipv4或ipv6 - ip/ipv6:    - null：客户端IP   - x-forwarded-for：请求header中X-Forwarded-For记录的IP   - $remote_addr：TCP连接IP   - $remote_sockaddr：3层源IP - params/cookie/header：   - check_all_indexes_logic为null： 可自定义子字段名称   - check_all_indexes_logic不为null：必须为null - 其他字段类型：不支持，默认为null  **默认取值：** 不涉及
	Index *string `json:"index,omitempty"`

	// **参数解释：** 匹配逻辑 **约束限制：** 匹配逻辑根据字段类型变化 **取值范围：** - url/user-agent/referer:    - contain_any: 包含任意一个   - not_contain_all: 不包含全部   - equal_any: 等于任意一个   - not_equal_all: 不等于全部   - prefix_any: 前缀匹配任意一个   - not_prefix_all: 前缀不匹配全部   - suffix_any: 后缀匹配任意一个   - not_suffix_all: 后缀不匹配全部   - contain: 包含   - not_contain: 不包含   - equal: 等于   - not_equal: 不等于   - prefix: 前缀匹配   - not_prefix: 前缀不匹配   - suffix: 后缀匹配   - not_suffix: 后缀不匹配   - len_equal: 长度等于   - len_not_equal: 长度不等于   - len_greater: 长度大于   - len_less: 长度小于   - len_greater_equal: 长度大于等于   - len_less_equal: 长度小于等于   - regular_match: 正则匹配   - regular_not_match: 正则不匹配 - custom_geoip:    - belong: 属于   - not_belong: 不属于 - robot:    - match: 匹配   - not_match: 不匹配 - ip/ipv6:   - equal_any: 等于任意一个   - not_equal_all: 不等于全部   - equal: 等于   - not_equal: 不等于 - params/cookie/header:   - contain_any: 包含任意一个   - not_contain_all: 不包含全部   - equal_any: 等于任意一个   - not_equal_all: 不等于全部   - prefix_any: 前缀匹配任意一个   - not_prefix_all: 前缀不匹配全部   - suffix_any: 后缀匹配任意一个   - not_suffix_all: 后缀不匹配全部   - contain: 包含   - not_contain: 不包含   - equal: 等于   - not_equal: 不等于   - prefix: 前缀匹配   - not_prefix: 前缀不匹配   - suffix: 后缀匹配   - not_suffix: 后缀不匹配   - len_equal: 长度等于   - len_not_equal: 长度不等于   - len_greater: 长度大于   - len_less: 长度小于   - len_greater_equal: 长度大于等于   - len_less_equal: 长度小于等于   - num_equal: 数字等于   - num_not_equal: 数字不等于   - num_greater: 数字大于   - num_less: 数字小于   - exist: 存在   - not_exist: 不存在   - regular_match: 正则匹配   - regular_not_match: 正则不匹配 - method/protocol:   - equal: 等于   - not_equal: 不等于 - request_line:   - len_equal: 长度等于   - len_not_equal: 长度不等于   - len_greater: 长度大于   - len_less: 长度小于   - len_greater_equal: 长度大于等于   - len_less_equal: 长度小于等于 - request:   - len_equal: 长度等于   - len_not_equal: 长度不等于   - len_greater: 长度大于   - len_less: 长度小于   - len_greater_equal: 长度大于等于   - len_less_equal: 长度小于等于   - regular_match: 正则匹配   - regular_not_match: 正则不匹配 - request_body:   - contain: 包含   - contain_any: 包含任意一个   - not_contain: 不包含   - not_contain_all: 不包含全部   - regular_match: 正则匹配   - regular_not_match: 正则不匹配  **默认取值：** 不涉及
	LogicOperation CreateConditionLogicOperation `json:"logic_operation"`

	// **参数解释：** 条件列表逻辑匹配内容 **约束限制：** 当logic_operation参数不以any或者all结尾时，需要传该参数 **取值范围：** 匹配内容字符串长度范围：[1, 4096] 内容格式根据参数category和logic_operation变化： - logic_operation为数值比较类型：纯数字 - url: URL格式；仅支持单个匹配内容 - custom_geoip: 客户端IP所属国家或省份，多个位置以|分隔，比如：\"BJ|SH\" - robot: 已知特征爬虫列表，可选择多个   - crawler_engine：搜索引擎   - crawler_scanner：扫描器   - crawler_script：脚本工具   - crawler_other：其他爬虫 - ip: IPv4 - ipv6: IPv6 - referer: 例如：http://test.com - params：不包含& - user-agent/cookie/header/request_body: 无限制 - method: HTTP协议支持的method，字母大写 - protocol:    - http   - https  **默认取值：** 不涉及
	Contents *[]string `json:"contents,omitempty"`

	// **参数解释：** 需要检查所有子字段或检查任意子字段时传此参数 **约束限制：** 仅当category为params/cookie/header时支持 **取值范围：** - 1：所有子字段 - 2：任意子字段  **默认取值：** 不涉及
	CheckAllIndexesLogic *int32 `json:"check_all_indexes_logic,omitempty"`

	// **参数解释：** 引用表ID **约束限制：** 当logic_operation参数以any或者all结尾时，需要传该参数；引用表类型要与category类型保持一致 **取值范围：** 通过ListValueList接口获取引用表ID  **默认取值：** 不涉及
	ValueListId *string `json:"value_list_id,omitempty"`
}

func (o CreateCondition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCondition struct{}"
	}

	return strings.Join([]string{"CreateCondition", string(data)}, " ")
}

type CreateConditionCategory struct {
	value string
}

type CreateConditionCategoryEnum struct {
	URL          CreateConditionCategory
	CUSTOM_GEO   CreateConditionCategory
	ROBOT        CreateConditionCategory
	USER_AGENT   CreateConditionCategory
	IP           CreateConditionCategory
	IPV6         CreateConditionCategory
	PARAMS       CreateConditionCategory
	COOKIE       CreateConditionCategory
	REFERER      CreateConditionCategory
	HEADER       CreateConditionCategory
	METHOD       CreateConditionCategory
	REQUEST_LINE CreateConditionCategory
	REQUEST      CreateConditionCategory
	PROTOCOL     CreateConditionCategory
	REQUEST_BODY CreateConditionCategory
}

func GetCreateConditionCategoryEnum() CreateConditionCategoryEnum {
	return CreateConditionCategoryEnum{
		URL: CreateConditionCategory{
			value: "url",
		},
		CUSTOM_GEO: CreateConditionCategory{
			value: "custom_geo",
		},
		ROBOT: CreateConditionCategory{
			value: "robot",
		},
		USER_AGENT: CreateConditionCategory{
			value: "user-agent",
		},
		IP: CreateConditionCategory{
			value: "ip",
		},
		IPV6: CreateConditionCategory{
			value: "ipv6",
		},
		PARAMS: CreateConditionCategory{
			value: "params",
		},
		COOKIE: CreateConditionCategory{
			value: "cookie",
		},
		REFERER: CreateConditionCategory{
			value: "referer",
		},
		HEADER: CreateConditionCategory{
			value: "header",
		},
		METHOD: CreateConditionCategory{
			value: "method",
		},
		REQUEST_LINE: CreateConditionCategory{
			value: "request_line",
		},
		REQUEST: CreateConditionCategory{
			value: "request",
		},
		PROTOCOL: CreateConditionCategory{
			value: "protocol",
		},
		REQUEST_BODY: CreateConditionCategory{
			value: "request_body",
		},
	}
}

func (c CreateConditionCategory) Value() string {
	return c.value
}

func (c CreateConditionCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateConditionCategory) UnmarshalJSON(b []byte) error {
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

type CreateConditionLogicOperation struct {
	value string
}

type CreateConditionLogicOperationEnum struct {
	LEN_EQUAL         CreateConditionLogicOperation
	LEN_NOT_EQUAL     CreateConditionLogicOperation
	LEN_GREATER       CreateConditionLogicOperation
	LEN_LESS          CreateConditionLogicOperation
	LEN_GREATER_EQUAL CreateConditionLogicOperation
	LEN_LESS_EQUAL    CreateConditionLogicOperation
	REGULAR_MATCH     CreateConditionLogicOperation
	REGULAR_NOT_MATCH CreateConditionLogicOperation
	MATCH             CreateConditionLogicOperation
	NOT_MATCH         CreateConditionLogicOperation
	NUM_EQUAL         CreateConditionLogicOperation
	NUM_NOT_EQUAL     CreateConditionLogicOperation
	NUM_GREATER       CreateConditionLogicOperation
	NUM_LESS          CreateConditionLogicOperation
	EXIST             CreateConditionLogicOperation
	NOT_EXIST         CreateConditionLogicOperation
	EQUAL             CreateConditionLogicOperation
	NOT_EQUAL         CreateConditionLogicOperation
	EQUAL_ANY         CreateConditionLogicOperation
	NOT_EQUAL_ALL     CreateConditionLogicOperation
	PREFIX            CreateConditionLogicOperation
	PREFIX_ANY        CreateConditionLogicOperation
	NOT_PREFIX        CreateConditionLogicOperation
	NOT_PREFIX_ALL    CreateConditionLogicOperation
	SUFFIX            CreateConditionLogicOperation
	SUFFIX_ANY        CreateConditionLogicOperation
	NOT_SUFFIX        CreateConditionLogicOperation
	NOT_SUFFIX_ALL    CreateConditionLogicOperation
	CONTAIN           CreateConditionLogicOperation
	CONTAIN_ANY       CreateConditionLogicOperation
	NOT_CONTAIN       CreateConditionLogicOperation
	NOT_CONTAIN_ALL   CreateConditionLogicOperation
}

func GetCreateConditionLogicOperationEnum() CreateConditionLogicOperationEnum {
	return CreateConditionLogicOperationEnum{
		LEN_EQUAL: CreateConditionLogicOperation{
			value: "len_equal",
		},
		LEN_NOT_EQUAL: CreateConditionLogicOperation{
			value: "len_not_equal",
		},
		LEN_GREATER: CreateConditionLogicOperation{
			value: "len_greater",
		},
		LEN_LESS: CreateConditionLogicOperation{
			value: "len_less",
		},
		LEN_GREATER_EQUAL: CreateConditionLogicOperation{
			value: "len_greater_equal",
		},
		LEN_LESS_EQUAL: CreateConditionLogicOperation{
			value: "len_less_equal",
		},
		REGULAR_MATCH: CreateConditionLogicOperation{
			value: "regular_match",
		},
		REGULAR_NOT_MATCH: CreateConditionLogicOperation{
			value: "regular_not_match",
		},
		MATCH: CreateConditionLogicOperation{
			value: "match",
		},
		NOT_MATCH: CreateConditionLogicOperation{
			value: "not_match",
		},
		NUM_EQUAL: CreateConditionLogicOperation{
			value: "num_equal",
		},
		NUM_NOT_EQUAL: CreateConditionLogicOperation{
			value: "num_not_equal",
		},
		NUM_GREATER: CreateConditionLogicOperation{
			value: "num_greater",
		},
		NUM_LESS: CreateConditionLogicOperation{
			value: "num_less",
		},
		EXIST: CreateConditionLogicOperation{
			value: "exist",
		},
		NOT_EXIST: CreateConditionLogicOperation{
			value: "not_exist",
		},
		EQUAL: CreateConditionLogicOperation{
			value: "equal",
		},
		NOT_EQUAL: CreateConditionLogicOperation{
			value: "not_equal",
		},
		EQUAL_ANY: CreateConditionLogicOperation{
			value: "equal_any",
		},
		NOT_EQUAL_ALL: CreateConditionLogicOperation{
			value: "not_equal_all",
		},
		PREFIX: CreateConditionLogicOperation{
			value: "prefix",
		},
		PREFIX_ANY: CreateConditionLogicOperation{
			value: "prefix_any",
		},
		NOT_PREFIX: CreateConditionLogicOperation{
			value: "not_prefix",
		},
		NOT_PREFIX_ALL: CreateConditionLogicOperation{
			value: "not_prefix_all",
		},
		SUFFIX: CreateConditionLogicOperation{
			value: "suffix",
		},
		SUFFIX_ANY: CreateConditionLogicOperation{
			value: "suffix_any",
		},
		NOT_SUFFIX: CreateConditionLogicOperation{
			value: "not_suffix",
		},
		NOT_SUFFIX_ALL: CreateConditionLogicOperation{
			value: "not_suffix_all",
		},
		CONTAIN: CreateConditionLogicOperation{
			value: "contain",
		},
		CONTAIN_ANY: CreateConditionLogicOperation{
			value: "contain_any",
		},
		NOT_CONTAIN: CreateConditionLogicOperation{
			value: "not_contain",
		},
		NOT_CONTAIN_ALL: CreateConditionLogicOperation{
			value: "not_contain_all",
		},
	}
}

func (c CreateConditionLogicOperation) Value() string {
	return c.value
}

func (c CreateConditionLogicOperation) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateConditionLogicOperation) UnmarshalJSON(b []byte) error {
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
