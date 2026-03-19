package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AnticrawlerCondition struct {

	// **参数解释：** 字段类型 **约束限制：** 不涉及 **取值范围：**  - url  - user-agent **默认取值：** 不涉及
	Category AnticrawlerConditionCategory `json:"category"`

	// **参数解释：** 条件列表匹配逻辑 **约束限制：** 不涉及 **取值范围：** - contain_any: 包含任意一个 - not_contain_all: 不包含全部 - equal_any: 等于任意一个 - not_equal_all: 不等于全部 - prefix_any: 前缀匹配任意一个 - not_prefix_all: 前缀不匹配全部 - suffix_any: 后缀匹配任意一个 - not_suffix_all: 后缀不匹配全部 - contain: 包含 - not_contain: 不包含 - equal: 等于 - not_equal: 不等于 - prefix: 前缀匹配 - not_prefix: 前缀不匹配 - suffix: 后缀匹配 - not_suffix: 后缀不匹配 - len_equal: 长度等于 - len_not_equal: 长度不等于 - len_greater: 长度大于 - len_less: 长度小于 - len_greater_equal: 长度大于等于 - len_less_equal: 长度小于等于 - regular_match: 正则匹配 - regular_not_match: 正则不匹配 **默认取值：** 不涉及
	LogicOperation AnticrawlerConditionLogicOperation `json:"logic_operation"`

	// **参数解释：** 条件列表逻辑匹配内容 **约束限制：** 当logic_operation参数不以any或者all结尾时，需要传该参数；仅支持单个匹配内容 **取值范围：** 匹配内容字符串长度范围：[1, 4096] **默认取值：** 不涉及
	Contents *[]string `json:"contents,omitempty"`

	// **参数解释：** 引用表ID **约束限制：** 当logic_operation参数以any或者all结尾时，需要传该参数；引用表类型要与category类型保持一致 **取值范围：** 通过ListValueList接口获取引用表ID **默认取值：** 不涉及
	ValueListId *string `json:"value_list_id,omitempty"`
}

func (o AnticrawlerCondition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AnticrawlerCondition struct{}"
	}

	return strings.Join([]string{"AnticrawlerCondition", string(data)}, " ")
}

type AnticrawlerConditionCategory struct {
	value string
}

type AnticrawlerConditionCategoryEnum struct {
	URL        AnticrawlerConditionCategory
	USER_AGENT AnticrawlerConditionCategory
}

func GetAnticrawlerConditionCategoryEnum() AnticrawlerConditionCategoryEnum {
	return AnticrawlerConditionCategoryEnum{
		URL: AnticrawlerConditionCategory{
			value: "url",
		},
		USER_AGENT: AnticrawlerConditionCategory{
			value: "user-agent",
		},
	}
}

func (c AnticrawlerConditionCategory) Value() string {
	return c.value
}

func (c AnticrawlerConditionCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AnticrawlerConditionCategory) UnmarshalJSON(b []byte) error {
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

type AnticrawlerConditionLogicOperation struct {
	value string
}

type AnticrawlerConditionLogicOperationEnum struct {
	CONTAIN_ANY       AnticrawlerConditionLogicOperation
	NOT_CONTAIN_ALL   AnticrawlerConditionLogicOperation
	EQUAL_ANY         AnticrawlerConditionLogicOperation
	NOT_EQUAL_ALL     AnticrawlerConditionLogicOperation
	PREFIX_ANY        AnticrawlerConditionLogicOperation
	NOT_PREFIX_ALL    AnticrawlerConditionLogicOperation
	SUFFIX_ANY        AnticrawlerConditionLogicOperation
	NOT_SUFFIX_ALL    AnticrawlerConditionLogicOperation
	CONTAIN           AnticrawlerConditionLogicOperation
	NOT_CONTAIN       AnticrawlerConditionLogicOperation
	EQUAL             AnticrawlerConditionLogicOperation
	NOT_EQUAL         AnticrawlerConditionLogicOperation
	PREFIX            AnticrawlerConditionLogicOperation
	NOT_PREFIX        AnticrawlerConditionLogicOperation
	SUFFIX            AnticrawlerConditionLogicOperation
	NOT_SUFFIX        AnticrawlerConditionLogicOperation
	LEN_EQUAL         AnticrawlerConditionLogicOperation
	LEN_NOT_EQUAL     AnticrawlerConditionLogicOperation
	LEN_GREATER       AnticrawlerConditionLogicOperation
	LEN_LESS          AnticrawlerConditionLogicOperation
	LEN_GREATER_EQUAL AnticrawlerConditionLogicOperation
	LEN_LESS_EQUAL    AnticrawlerConditionLogicOperation
	REGULAR_MATCH     AnticrawlerConditionLogicOperation
	REGULAR_NOT_MATCH AnticrawlerConditionLogicOperation
}

func GetAnticrawlerConditionLogicOperationEnum() AnticrawlerConditionLogicOperationEnum {
	return AnticrawlerConditionLogicOperationEnum{
		CONTAIN_ANY: AnticrawlerConditionLogicOperation{
			value: "contain_any",
		},
		NOT_CONTAIN_ALL: AnticrawlerConditionLogicOperation{
			value: "not_contain_all",
		},
		EQUAL_ANY: AnticrawlerConditionLogicOperation{
			value: "equal_any",
		},
		NOT_EQUAL_ALL: AnticrawlerConditionLogicOperation{
			value: "not_equal_all",
		},
		PREFIX_ANY: AnticrawlerConditionLogicOperation{
			value: "prefix_any",
		},
		NOT_PREFIX_ALL: AnticrawlerConditionLogicOperation{
			value: "not_prefix_all",
		},
		SUFFIX_ANY: AnticrawlerConditionLogicOperation{
			value: "suffix_any",
		},
		NOT_SUFFIX_ALL: AnticrawlerConditionLogicOperation{
			value: "not_suffix_all",
		},
		CONTAIN: AnticrawlerConditionLogicOperation{
			value: "contain",
		},
		NOT_CONTAIN: AnticrawlerConditionLogicOperation{
			value: "not_contain",
		},
		EQUAL: AnticrawlerConditionLogicOperation{
			value: "equal",
		},
		NOT_EQUAL: AnticrawlerConditionLogicOperation{
			value: "not_equal",
		},
		PREFIX: AnticrawlerConditionLogicOperation{
			value: "prefix",
		},
		NOT_PREFIX: AnticrawlerConditionLogicOperation{
			value: "not_prefix",
		},
		SUFFIX: AnticrawlerConditionLogicOperation{
			value: "suffix",
		},
		NOT_SUFFIX: AnticrawlerConditionLogicOperation{
			value: "not_suffix",
		},
		LEN_EQUAL: AnticrawlerConditionLogicOperation{
			value: "len_equal",
		},
		LEN_NOT_EQUAL: AnticrawlerConditionLogicOperation{
			value: "len_not_equal",
		},
		LEN_GREATER: AnticrawlerConditionLogicOperation{
			value: "len_greater",
		},
		LEN_LESS: AnticrawlerConditionLogicOperation{
			value: "len_less",
		},
		LEN_GREATER_EQUAL: AnticrawlerConditionLogicOperation{
			value: "len_greater_equal",
		},
		LEN_LESS_EQUAL: AnticrawlerConditionLogicOperation{
			value: "len_less_equal",
		},
		REGULAR_MATCH: AnticrawlerConditionLogicOperation{
			value: "regular_match",
		},
		REGULAR_NOT_MATCH: AnticrawlerConditionLogicOperation{
			value: "regular_not_match",
		},
	}
}

func (c AnticrawlerConditionLogicOperation) Value() string {
	return c.value
}

func (c AnticrawlerConditionLogicOperation) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AnticrawlerConditionLogicOperation) UnmarshalJSON(b []byte) error {
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
