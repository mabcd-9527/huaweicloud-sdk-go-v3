package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowLockBlockingStatisticsRequest Request Object
type ShowLockBlockingStatisticsRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 当前时间戳（Unix timestamp），单位：毫秒。
	CurrentTime int64 `json:"current_time"`

	// 请求语言类型。en-us：英文。 zh-cn：中文。
	XLanguage *ShowLockBlockingStatisticsRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ShowLockBlockingStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLockBlockingStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ShowLockBlockingStatisticsRequest", string(data)}, " ")
}

type ShowLockBlockingStatisticsRequestXLanguage struct {
	value string
}

type ShowLockBlockingStatisticsRequestXLanguageEnum struct {
	ZH_CN ShowLockBlockingStatisticsRequestXLanguage
	EN_US ShowLockBlockingStatisticsRequestXLanguage
}

func GetShowLockBlockingStatisticsRequestXLanguageEnum() ShowLockBlockingStatisticsRequestXLanguageEnum {
	return ShowLockBlockingStatisticsRequestXLanguageEnum{
		ZH_CN: ShowLockBlockingStatisticsRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowLockBlockingStatisticsRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowLockBlockingStatisticsRequestXLanguage) Value() string {
	return c.value
}

func (c ShowLockBlockingStatisticsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowLockBlockingStatisticsRequestXLanguage) UnmarshalJSON(b []byte) error {
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
