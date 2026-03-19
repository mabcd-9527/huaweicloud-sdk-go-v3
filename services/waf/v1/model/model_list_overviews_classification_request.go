package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListOverviewsClassificationRequest Request Object
type ListOverviewsClassificationRequest struct {

	// **参数解释：** 地理位置展示语言 **约束限制：** 不涉及 **取值范围：** - zh-cn 中文 - en-us 英文 **默认取值：** en-us
	XLanguage *ListOverviewsClassificationRequestXLanguage `json:"X-Language,omitempty"`

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id。若需要查询当前用户所有企业项目绑定的资源信息，请传参all_granted_eps。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释：** 起始时间(毫秒时间戳)，需要和to同时使用 **约束限制：** from <= to **取值范围：** from ~ to 最大范围30天 **默认取值：** 不涉及
	From int64 `json:"from"`

	// **参数解释：** 结束时间(毫秒时间戳)，需要和from同时使用 **约束限制：** from ~ to 最大范围30天 **取值范围：** 不能超过当天的结束时间 **默认取值：** 不涉及
	To int64 `json:"to"`

	// **参数解释：** 查询前TopN的结果 **约束限制：** 不涉及 **取值范围：** [1, 10] **默认取值：** 5
	Top *int32 `json:"top,omitempty"`

	// **参数解释：** 要查询的域名id列表，通过 ”查询独享模式域名列表“（ListPremiumHost）或者 “查询云模式防护域名列表” （ListHost）接口获取；不传参代表查询全部域名的数据 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Hosts *[]string `json:"hosts,omitempty"`

	// **参数解释：** 要查询的实例id列表，通过 “查询WAF独享引擎列表”（ListInstance）接口获取 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Instances *[]string `json:"instances,omitempty"`
}

func (o ListOverviewsClassificationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOverviewsClassificationRequest struct{}"
	}

	return strings.Join([]string{"ListOverviewsClassificationRequest", string(data)}, " ")
}

type ListOverviewsClassificationRequestXLanguage struct {
	value string
}

type ListOverviewsClassificationRequestXLanguageEnum struct {
	ZH_CN ListOverviewsClassificationRequestXLanguage
	EN_US ListOverviewsClassificationRequestXLanguage
}

func GetListOverviewsClassificationRequestXLanguageEnum() ListOverviewsClassificationRequestXLanguageEnum {
	return ListOverviewsClassificationRequestXLanguageEnum{
		ZH_CN: ListOverviewsClassificationRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListOverviewsClassificationRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListOverviewsClassificationRequestXLanguage) Value() string {
	return c.value
}

func (c ListOverviewsClassificationRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListOverviewsClassificationRequestXLanguage) UnmarshalJSON(b []byte) error {
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
