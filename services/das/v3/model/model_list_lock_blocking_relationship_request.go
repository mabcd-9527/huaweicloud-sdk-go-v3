package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListLockBlockingRelationshipRequest Request Object
type ListLockBlockingRelationshipRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 批次ID。
	UniqueId string `json:"unique_id"`

	// 会话ID。
	Spid int64 `json:"spid"`

	// 请求语言类型。en-us：英文。 zh-cn：中文。
	XLanguage *ListLockBlockingRelationshipRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ListLockBlockingRelationshipRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLockBlockingRelationshipRequest struct{}"
	}

	return strings.Join([]string{"ListLockBlockingRelationshipRequest", string(data)}, " ")
}

type ListLockBlockingRelationshipRequestXLanguage struct {
	value string
}

type ListLockBlockingRelationshipRequestXLanguageEnum struct {
	ZH_CN ListLockBlockingRelationshipRequestXLanguage
	EN_US ListLockBlockingRelationshipRequestXLanguage
}

func GetListLockBlockingRelationshipRequestXLanguageEnum() ListLockBlockingRelationshipRequestXLanguageEnum {
	return ListLockBlockingRelationshipRequestXLanguageEnum{
		ZH_CN: ListLockBlockingRelationshipRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListLockBlockingRelationshipRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListLockBlockingRelationshipRequestXLanguage) Value() string {
	return c.value
}

func (c ListLockBlockingRelationshipRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListLockBlockingRelationshipRequestXLanguage) UnmarshalJSON(b []byte) error {
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
