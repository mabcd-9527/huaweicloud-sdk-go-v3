package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListLockBlockingDetailRequest Request Object
type ListLockBlockingDetailRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 开始时间戳（Unix timestamp），单位：毫秒。
	StartTime int64 `json:"start_time"`

	// 结束时间戳（Unix timestamp），单位：毫秒。
	EndTime int64 `json:"end_time"`

	// 每页返回的数量，默认值为20
	PerPage *int32 `json:"per_page,omitempty"`

	// 当前页码，第一次查询的时候默认值为1
	CurPage *int32 `json:"cur_page,omitempty"`

	// 数据库名。
	DbName *string `json:"db_name,omitempty"`

	// 请求语言类型。en-us：英文。 zh-cn：中文。
	XLanguage *ListLockBlockingDetailRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ListLockBlockingDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLockBlockingDetailRequest struct{}"
	}

	return strings.Join([]string{"ListLockBlockingDetailRequest", string(data)}, " ")
}

type ListLockBlockingDetailRequestXLanguage struct {
	value string
}

type ListLockBlockingDetailRequestXLanguageEnum struct {
	ZH_CN ListLockBlockingDetailRequestXLanguage
	EN_US ListLockBlockingDetailRequestXLanguage
}

func GetListLockBlockingDetailRequestXLanguageEnum() ListLockBlockingDetailRequestXLanguageEnum {
	return ListLockBlockingDetailRequestXLanguageEnum{
		ZH_CN: ListLockBlockingDetailRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListLockBlockingDetailRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListLockBlockingDetailRequestXLanguage) Value() string {
	return c.value
}

func (c ListLockBlockingDetailRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListLockBlockingDetailRequestXLanguage) UnmarshalJSON(b []byte) error {
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
