package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowInstanceBackupsRequest Request Object
type ShowInstanceBackupsRequest struct {

	// **参数解释**：  请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us - zh-cn  **默认取值**：  en-us。
	XLanguage *ShowInstanceBackupsRequestXLanguage `json:"X-Language,omitempty"`

	// **参数解释**：  实例ID，此参数是实例的唯一标识。  获取方法请参见[查询实例列表](https://support.huaweicloud.com/api-taurusdb/ListGaussMySqlInstancesUnifyStatus.html)。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，后缀为in07，长度为36个字符。  **默认取值**：  不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**：    索引位置，偏移量。从第一条数据偏移offset条数据后开始查询。    **约束限制**：    必须为整数，不能为负数。    **取值范围**：    ≥0。  **默认取值**：    0。
	Offset *string `json:"offset,omitempty"`

	// **参数解释**：  查询记录数。  **约束限制**：  必须为整数，不能为负数。  **取值范围**：  1-100。  **默认取值**：  100。
	Limit *string `json:"limit,omitempty"`

	// **参数解释**:  根据指定字段排序。  **约束限制**:  不涉及。  **取值范围**:  - name：备份名称。 - beginTime：备份开启时间。 - type：备份类型。  **默认取值**: 不涉及。
	OrderField *string `json:"order_field,omitempty"`

	// **参数解释**: 排序规则。  **约束限制**: 不涉及。  **取值范围**: - asc：升序。 - desc：降序。  **默认取值**: 不涉及。
	OrderRule *string `json:"order_rule,omitempty"`

	// **参数解释**: 过滤字段类型。 **约束限制**: 不涉及。 **取值范围**: name：根据备份名称进行过滤。 **默认取值**: 不涉及。
	FilterField *string `json:"filter_field,omitempty"`

	// **参数解释**: 过滤内容。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	FilterContent *string `json:"filter_content,omitempty"`
}

func (o ShowInstanceBackupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceBackupsRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceBackupsRequest", string(data)}, " ")
}

type ShowInstanceBackupsRequestXLanguage struct {
	value string
}

type ShowInstanceBackupsRequestXLanguageEnum struct {
	ZH_CN ShowInstanceBackupsRequestXLanguage
	EN_US ShowInstanceBackupsRequestXLanguage
}

func GetShowInstanceBackupsRequestXLanguageEnum() ShowInstanceBackupsRequestXLanguageEnum {
	return ShowInstanceBackupsRequestXLanguageEnum{
		ZH_CN: ShowInstanceBackupsRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowInstanceBackupsRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowInstanceBackupsRequestXLanguage) Value() string {
	return c.value
}

func (c ShowInstanceBackupsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowInstanceBackupsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
