package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowEventRequest Request Object
type ShowEventRequest struct {

	// **参数解释：** 客户端IP所属地理位置展示语言，默认值为en-us **约束限制：** 不涉及 **取值范围：** - zh-cn 中文 - en-us 英文 **默认取值：** en-us
	XLanguage *ShowEventRequestXLanguage `json:"X-Language,omitempty"`

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id。若需要查询当前用户所有企业项目绑定的资源信息，请传参all_granted_eps。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释：** 防护事件id,通过调用查询攻击事件列表(ListEvent)接口获取防护事件id **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Eventid string `json:"eventid"`
}

func (o ShowEventRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEventRequest struct{}"
	}

	return strings.Join([]string{"ShowEventRequest", string(data)}, " ")
}

type ShowEventRequestXLanguage struct {
	value string
}

type ShowEventRequestXLanguageEnum struct {
	ZH_CN ShowEventRequestXLanguage
	EN_US ShowEventRequestXLanguage
}

func GetShowEventRequestXLanguageEnum() ShowEventRequestXLanguageEnum {
	return ShowEventRequestXLanguageEnum{
		ZH_CN: ShowEventRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowEventRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowEventRequestXLanguage) Value() string {
	return c.value
}

func (c ShowEventRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowEventRequestXLanguage) UnmarshalJSON(b []byte) error {
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
