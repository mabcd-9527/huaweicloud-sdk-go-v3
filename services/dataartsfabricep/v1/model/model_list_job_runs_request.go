package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListJobRunsRequest Request Object
type ListJobRunsRequest struct {

	// Workspace的ID
	WorkspaceId string `json:"workspace_id"`

	// 通过作业运行id检索
	Id *string `json:"id,omitempty"`

	// 通过名字搜索运行的作业
	Name *string `json:"name,omitempty"`

	// 指定每一页返回的最大条目数，取值范围[1,100]，默认为10。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，表示从此偏移量开始查询， offset大于等于0，默认为0。
	Offset *int32 `json:"offset,omitempty"`

	// 通过作业id检索
	JobId *string `json:"job_id,omitempty"`

	// 通过id检索Endpoint的参数
	EndpointId *string `json:"endpoint_id,omitempty"`

	// 通过作业版本检索
	VersionId *string `json:"version_id,omitempty"`

	// 状态过滤，支持一种状态查询，默认查询所有；  可选值有： PENDING,CREATING,RUNNING,UPDATING,SUCCEEDED,FAILED,STOPPING,STOPPED,DELETING,DELETED
	Status *string `json:"status,omitempty"`

	// 类别，可选值如下：  - RAY_JOB: Ray Job,   - PYTHON_JOB: Python Job,   - SPARK_JOB: Spark Job
	Type *ListJobRunsRequestType `json:"type,omitempty"`

	// 根据字段排序，可选值： - CREATE_TIME：创建时间 - DURATION: 运行时间
	SortBy *ListJobRunsRequestSortBy `json:"sort_by,omitempty"`

	// 排序方式，可选值： - ASC：正序排序 - DESC: 倒序排序
	OrderBy *ListJobRunsRequestOrderBy `json:"order_by,omitempty"`
}

func (o ListJobRunsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobRunsRequest struct{}"
	}

	return strings.Join([]string{"ListJobRunsRequest", string(data)}, " ")
}

type ListJobRunsRequestType struct {
	value string
}

type ListJobRunsRequestTypeEnum struct {
	RAY_JOB    ListJobRunsRequestType
	PYTHON_JOB ListJobRunsRequestType
	SPARK_JOB  ListJobRunsRequestType
}

func GetListJobRunsRequestTypeEnum() ListJobRunsRequestTypeEnum {
	return ListJobRunsRequestTypeEnum{
		RAY_JOB: ListJobRunsRequestType{
			value: "RAY_JOB",
		},
		PYTHON_JOB: ListJobRunsRequestType{
			value: "PYTHON_JOB",
		},
		SPARK_JOB: ListJobRunsRequestType{
			value: "SPARK_JOB",
		},
	}
}

func (c ListJobRunsRequestType) Value() string {
	return c.value
}

func (c ListJobRunsRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListJobRunsRequestType) UnmarshalJSON(b []byte) error {
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

type ListJobRunsRequestSortBy struct {
	value string
}

type ListJobRunsRequestSortByEnum struct {
	CREATE_TIME ListJobRunsRequestSortBy
	DURATION    ListJobRunsRequestSortBy
}

func GetListJobRunsRequestSortByEnum() ListJobRunsRequestSortByEnum {
	return ListJobRunsRequestSortByEnum{
		CREATE_TIME: ListJobRunsRequestSortBy{
			value: "CREATE_TIME",
		},
		DURATION: ListJobRunsRequestSortBy{
			value: "DURATION",
		},
	}
}

func (c ListJobRunsRequestSortBy) Value() string {
	return c.value
}

func (c ListJobRunsRequestSortBy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListJobRunsRequestSortBy) UnmarshalJSON(b []byte) error {
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

type ListJobRunsRequestOrderBy struct {
	value string
}

type ListJobRunsRequestOrderByEnum struct {
	ASC  ListJobRunsRequestOrderBy
	DESC ListJobRunsRequestOrderBy
}

func GetListJobRunsRequestOrderByEnum() ListJobRunsRequestOrderByEnum {
	return ListJobRunsRequestOrderByEnum{
		ASC: ListJobRunsRequestOrderBy{
			value: "ASC",
		},
		DESC: ListJobRunsRequestOrderBy{
			value: "DESC",
		},
	}
}

func (c ListJobRunsRequestOrderBy) Value() string {
	return c.value
}

func (c ListJobRunsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListJobRunsRequestOrderBy) UnmarshalJSON(b []byte) error {
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
