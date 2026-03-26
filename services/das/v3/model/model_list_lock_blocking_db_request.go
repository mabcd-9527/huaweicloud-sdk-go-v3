package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListLockBlockingDbRequest Request Object
type ListLockBlockingDbRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 开始时间戳（Unix timestamp），单位：毫秒。
	StartTime int64 `json:"start_time"`

	// 结束时间戳（Unix timestamp），单位：毫秒。
	EndTime int64 `json:"end_time"`

	// 请求语言类型。en-us：英文。 zh-cn：中文。
	XLanguage *ListLockBlockingDbRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ListLockBlockingDbRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLockBlockingDbRequest struct{}"
	}

	return strings.Join([]string{"ListLockBlockingDbRequest", string(data)}, " ")
}

type ListLockBlockingDbRequestXLanguage struct {
	value string
}

type ListLockBlockingDbRequestXLanguageEnum struct {
	ZH_CN ListLockBlockingDbRequestXLanguage
	EN_US ListLockBlockingDbRequestXLanguage
}

func GetListLockBlockingDbRequestXLanguageEnum() ListLockBlockingDbRequestXLanguageEnum {
	return ListLockBlockingDbRequestXLanguageEnum{
		ZH_CN: ListLockBlockingDbRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListLockBlockingDbRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListLockBlockingDbRequestXLanguage) Value() string {
	return c.value
}

func (c ListLockBlockingDbRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListLockBlockingDbRequestXLanguage) UnmarshalJSON(b []byte) error {
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
