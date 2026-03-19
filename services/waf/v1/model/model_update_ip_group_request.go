package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateIpGroupRequest Request Object
type UpdateIpGroupRequest struct {

	// **参数解释：** 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目ID。若需要查询当前用户所有企业项目绑定的资源信息，请传参all_granted_eps。 **约束限制：** 不涉及 **取值范围：**  - 0：代表default企业项目  - all_granted_eps：代表所有企业项目  - 其它企业项目ID：长度为36个字符 **默认取值：** 0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释：** ip地址组id，可从查询地址组列表(ListIpGroup)接口中获取 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Id string `json:"id"`

	// **参数解释：** 修改ip地址组时，此为必传字段 **约束限制：** 不涉及 **取值范围：** - add 添加 - delete 删除 - update 修改 **默认取值：** 不涉及
	Action *UpdateIpGroupRequestAction `json:"action,omitempty"`

	Body *UpdateIpGroupRequestBody `json:"body,omitempty"`
}

func (o UpdateIpGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIpGroupRequest struct{}"
	}

	return strings.Join([]string{"UpdateIpGroupRequest", string(data)}, " ")
}

type UpdateIpGroupRequestAction struct {
	value string
}

type UpdateIpGroupRequestActionEnum struct {
	ADD    UpdateIpGroupRequestAction
	DELETE UpdateIpGroupRequestAction
	UPDATE UpdateIpGroupRequestAction
}

func GetUpdateIpGroupRequestActionEnum() UpdateIpGroupRequestActionEnum {
	return UpdateIpGroupRequestActionEnum{
		ADD: UpdateIpGroupRequestAction{
			value: "add",
		},
		DELETE: UpdateIpGroupRequestAction{
			value: "delete",
		},
		UPDATE: UpdateIpGroupRequestAction{
			value: "update",
		},
	}
}

func (c UpdateIpGroupRequestAction) Value() string {
	return c.value
}

func (c UpdateIpGroupRequestAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateIpGroupRequestAction) UnmarshalJSON(b []byte) error {
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
