package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSourceIpTop5Request Request Object
type ListSourceIpTop5Request struct {

	// **参数解释：** 查询的时间范围，recent参数与from、to必须使用其中一个。当同时使用recent参数与from、to时，以recent参数为准 **约束限制：** 不涉及 **取值范围：**  - yesterday：昨天  - today：今天  - 3days：近3天   - 1week：近7天   - 1month：近30天  **默认取值：** 不涉及
	Recent *ListSourceIpTop5RequestRecent `json:"recent,omitempty"`

	// **参数解释：** 起始时间(毫秒时间戳)，需要和to同时使用 **约束限制：** from <= to **取值范围：** from ~ to 最大范围30天 **默认取值：** 不涉及
	From *int32 `json:"from,omitempty"`

	// **参数解释：** 结束时间(毫秒时间戳)，需要和from同时使用 **约束限制：** from ~ to 最大范围30天 **取值范围：** 不能超过当天的结束时间 **默认取值：** 不涉及
	To *int32 `json:"to,omitempty"`

	// **参数解释：** 查询前TopN的结果 **约束限制：** 不涉及 **取值范围：** [1, 10] **默认取值：** 5
	Top *int32 `json:"top,omitempty"`

	// **参数解释：** 要查询的域名id列表，通过 ”查询独享模式域名列表“（ListPremiumHost）或者 “查询云模式防护域名列表” （ListHost）接口获取；不传参代表查询全部域名的数据 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Hosts *[]string `json:"hosts,omitempty"`

	// **参数解释：** 要查询的实例id列表，通过 “查询WAF独享引擎列表”（ListInstance）接口获取 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Instances *[]string `json:"instances,omitempty"`
}

func (o ListSourceIpTop5Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSourceIpTop5Request struct{}"
	}

	return strings.Join([]string{"ListSourceIpTop5Request", string(data)}, " ")
}

type ListSourceIpTop5RequestRecent struct {
	value string
}

type ListSourceIpTop5RequestRecentEnum struct {
	YESTERDAY ListSourceIpTop5RequestRecent
	TODAY     ListSourceIpTop5RequestRecent
	E_3DAYS   ListSourceIpTop5RequestRecent
	E_1WEEK   ListSourceIpTop5RequestRecent
	E_1MONTH  ListSourceIpTop5RequestRecent
}

func GetListSourceIpTop5RequestRecentEnum() ListSourceIpTop5RequestRecentEnum {
	return ListSourceIpTop5RequestRecentEnum{
		YESTERDAY: ListSourceIpTop5RequestRecent{
			value: "yesterday",
		},
		TODAY: ListSourceIpTop5RequestRecent{
			value: "today",
		},
		E_3DAYS: ListSourceIpTop5RequestRecent{
			value: "3days",
		},
		E_1WEEK: ListSourceIpTop5RequestRecent{
			value: "1week",
		},
		E_1MONTH: ListSourceIpTop5RequestRecent{
			value: "1month",
		},
	}
}

func (c ListSourceIpTop5RequestRecent) Value() string {
	return c.value
}

func (c ListSourceIpTop5RequestRecent) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSourceIpTop5RequestRecent) UnmarshalJSON(b []byte) error {
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
