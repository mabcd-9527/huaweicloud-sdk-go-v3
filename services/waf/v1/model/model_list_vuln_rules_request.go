package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListVulnRulesRequest Request Object
type ListVulnRulesRequest struct {

	// **参数解释：** 语言类型 **约束限制：** 不涉及 **取值范围：** - zh-cn 中文 - en-us 英文 **默认取值：** zh-cn
	XLanguage *ListVulnRulesRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释：** 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目ID。若需要查询当前用户所有企业项目绑定的资源信息，请传参all_granted_eps。 **约束限制：** 不涉及 **取值范围：**  - 0：代表default企业项目  - all_granted_eps：代表所有企业项目  - 其它企业项目ID：长度为36个字符 **默认取值：** 0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释：** 分页查询的起始位置，表示从第几条记录开始返回 **约束限制：** 不涉及 **取值范围：** [0,2147483645] **默认取值：** 0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 分页查询的单页返回数量，控制每次请求返回的记录条数。 **约束限制：** 不涉及 **取值范围：** [1, 1000] **默认取值：** 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：** 起始时间（13位毫秒时间戳），需要和to同时使用。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	From *int64 `json:"from,omitempty"`

	// **参数解释：** 结束时间（13位毫秒时间戳），需要和from同时使用。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	To *int64 `json:"to,omitempty"`
}

func (o ListVulnRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVulnRulesRequest struct{}"
	}

	return strings.Join([]string{"ListVulnRulesRequest", string(data)}, " ")
}

type ListVulnRulesRequestXLanguage struct {
	value string
}

type ListVulnRulesRequestXLanguageEnum struct {
	ZH_CN ListVulnRulesRequestXLanguage
	EN_US ListVulnRulesRequestXLanguage
}

func GetListVulnRulesRequestXLanguageEnum() ListVulnRulesRequestXLanguageEnum {
	return ListVulnRulesRequestXLanguageEnum{
		ZH_CN: ListVulnRulesRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListVulnRulesRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListVulnRulesRequestXLanguage) Value() string {
	return c.value
}

func (c ListVulnRulesRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListVulnRulesRequestXLanguage) UnmarshalJSON(b []byte) error {
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
