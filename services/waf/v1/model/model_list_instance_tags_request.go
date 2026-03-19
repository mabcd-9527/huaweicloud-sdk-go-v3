package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListInstanceTagsRequest Request Object
type ListInstanceTagsRequest struct {

	// **参数解释：** 资源类型 **约束限制：** 不涉及 **取值范围：** - waf 云模式引擎 - waf-instance 独享引擎 **默认取值：** 不涉及
	ResourceType ListInstanceTagsRequestResourceType `json:"resource_type"`

	// **参数解释：** 引擎id **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Resourceid string `json:"resourceid"`

	// **参数解释：** 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目ID。若需要查询当前用户所有企业项目绑定的资源信息，请传参all_granted_eps。 **约束限制：** 不涉及 **取值范围：**  - 0：代表default企业项目  - all_granted_eps：代表所有企业项目  - 其它企业项目ID：长度为36个字符 **默认取值：** 0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ListInstanceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceTagsRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceTagsRequest", string(data)}, " ")
}

type ListInstanceTagsRequestResourceType struct {
	value string
}

type ListInstanceTagsRequestResourceTypeEnum struct {
	WAF          ListInstanceTagsRequestResourceType
	WAF_INSTANCE ListInstanceTagsRequestResourceType
}

func GetListInstanceTagsRequestResourceTypeEnum() ListInstanceTagsRequestResourceTypeEnum {
	return ListInstanceTagsRequestResourceTypeEnum{
		WAF: ListInstanceTagsRequestResourceType{
			value: "waf",
		},
		WAF_INSTANCE: ListInstanceTagsRequestResourceType{
			value: "waf-instance",
		},
	}
}

func (c ListInstanceTagsRequestResourceType) Value() string {
	return c.value
}

func (c ListInstanceTagsRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstanceTagsRequestResourceType) UnmarshalJSON(b []byte) error {
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
