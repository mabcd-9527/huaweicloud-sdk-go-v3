package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowLockBlockingTrendRequest Request Object
type ShowLockBlockingTrendRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 开始时间戳（Unix timestamp），单位：毫秒。
	StartTime int64 `json:"start_time"`

	// 结束时间戳（Unix timestamp），单位：毫秒。
	EndTime int64 `json:"end_time"`

	// 请求语言类型。en-us：英文。 zh-cn：中文。
	XLanguage *ShowLockBlockingTrendRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ShowLockBlockingTrendRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLockBlockingTrendRequest struct{}"
	}

	return strings.Join([]string{"ShowLockBlockingTrendRequest", string(data)}, " ")
}

type ShowLockBlockingTrendRequestXLanguage struct {
	value string
}

type ShowLockBlockingTrendRequestXLanguageEnum struct {
	ZH_CN ShowLockBlockingTrendRequestXLanguage
	EN_US ShowLockBlockingTrendRequestXLanguage
}

func GetShowLockBlockingTrendRequestXLanguageEnum() ShowLockBlockingTrendRequestXLanguageEnum {
	return ShowLockBlockingTrendRequestXLanguageEnum{
		ZH_CN: ShowLockBlockingTrendRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowLockBlockingTrendRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowLockBlockingTrendRequestXLanguage) Value() string {
	return c.value
}

func (c ShowLockBlockingTrendRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowLockBlockingTrendRequestXLanguage) UnmarshalJSON(b []byte) error {
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
