package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreatePrivacyRuleRequestBody struct {

	// **参数解释：** 隐私屏蔽规则防护的url，需要填写标准的url格式，例如/admin/xxx或者/admin/\\*，以\"\\*\"结尾的路径前缀表示一个通配符，用于匹配该路径前缀下的所有子路径。例如，如果你有一个路径前缀/admin/，那么它将匹配所有以/admin/开头的URL路径。将url的参数设置为/admin/\\*，所有以/admin/开头的URL路径都会被该规则所覆盖 **约束限制：** 标准URL字符串或者URL前缀匹配字符串 **取值范围：** 不涉及 **默认取值：** 不涉及
	Url string `json:"url"`

	// **参数解释：** 屏蔽字段 **约束限制：** 不涉及 **取值范围：**  - params: 请求参数  - cookie: 根据Cookie区分的Web访问者  - header: 自定义HTTP首部  - form: 表单参数 **默认取值：** 不涉及
	Category CreatePrivacyRuleRequestBodyCategory `json:"category"`

	// **参数解释：** 屏蔽字段名，被屏蔽的字段将不会出现在日志中 **约束限制：** 长度不能超过2048字节 **取值范围：** 不涉及 **默认取值：** 不涉及
	Index string `json:"index"`

	// 规则描述，可选参数，设置该规则的备注信息。
	Description *string `json:"description,omitempty"`
}

func (o CreatePrivacyRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePrivacyRuleRequestBody struct{}"
	}

	return strings.Join([]string{"CreatePrivacyRuleRequestBody", string(data)}, " ")
}

type CreatePrivacyRuleRequestBodyCategory struct {
	value string
}

type CreatePrivacyRuleRequestBodyCategoryEnum struct {
	PARAMS CreatePrivacyRuleRequestBodyCategory
	COOKIE CreatePrivacyRuleRequestBodyCategory
	HEADER CreatePrivacyRuleRequestBodyCategory
	FORM   CreatePrivacyRuleRequestBodyCategory
}

func GetCreatePrivacyRuleRequestBodyCategoryEnum() CreatePrivacyRuleRequestBodyCategoryEnum {
	return CreatePrivacyRuleRequestBodyCategoryEnum{
		PARAMS: CreatePrivacyRuleRequestBodyCategory{
			value: "params",
		},
		COOKIE: CreatePrivacyRuleRequestBodyCategory{
			value: "cookie",
		},
		HEADER: CreatePrivacyRuleRequestBodyCategory{
			value: "header",
		},
		FORM: CreatePrivacyRuleRequestBodyCategory{
			value: "form",
		},
	}
}

func (c CreatePrivacyRuleRequestBodyCategory) Value() string {
	return c.value
}

func (c CreatePrivacyRuleRequestBodyCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreatePrivacyRuleRequestBodyCategory) UnmarshalJSON(b []byte) error {
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
