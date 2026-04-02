package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListGcTasksRequest Request Object
type ListGcTasksRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	// 分页查询时的查询标记，使用上一次接口调用返回的next_marker值，默认值从第一条数据查询。**注意：marker和limit参数需要配套使用。**
	Marker *string `json:"marker,omitempty"`

	// 条目数量，用于分页查询，默认值为10，最大值为100
	Limit *int32 `json:"limit,omitempty"`

	// 任务状态，Success：已完成，Stopped：已停止，Running：清理中，Pending：排队中，Error：失败。
	Status *ListGcTasksRequestStatus `json:"status,omitempty"`
}

func (o ListGcTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGcTasksRequest struct{}"
	}

	return strings.Join([]string{"ListGcTasksRequest", string(data)}, " ")
}

type ListGcTasksRequestStatus struct {
	value string
}

type ListGcTasksRequestStatusEnum struct {
	SUCCESS ListGcTasksRequestStatus
	STOPPED ListGcTasksRequestStatus
	RUNNING ListGcTasksRequestStatus
	PENDING ListGcTasksRequestStatus
	ERROR   ListGcTasksRequestStatus
}

func GetListGcTasksRequestStatusEnum() ListGcTasksRequestStatusEnum {
	return ListGcTasksRequestStatusEnum{
		SUCCESS: ListGcTasksRequestStatus{
			value: "Success",
		},
		STOPPED: ListGcTasksRequestStatus{
			value: "Stopped",
		},
		RUNNING: ListGcTasksRequestStatus{
			value: "Running",
		},
		PENDING: ListGcTasksRequestStatus{
			value: "Pending",
		},
		ERROR: ListGcTasksRequestStatus{
			value: "Error",
		},
	}
}

func (c ListGcTasksRequestStatus) Value() string {
	return c.value
}

func (c ListGcTasksRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListGcTasksRequestStatus) UnmarshalJSON(b []byte) error {
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
