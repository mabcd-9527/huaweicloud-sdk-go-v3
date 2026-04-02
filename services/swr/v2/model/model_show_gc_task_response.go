package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowGcTaskResponse Response Object
type ShowGcTaskResponse struct {

	// gc任务的ID。
	Id *int32 `json:"id,omitempty"`

	// gc任务的名称。
	JobName *string `json:"job_name,omitempty"`

	// 任务类型，MANUAL：手动执行，SCHEDULE：定时执行。
	JobKind *ShowGcTaskResponseJobKind `json:"job_kind,omitempty"`

	JobParameters *JobParameters `json:"job_parameters,omitempty"`

	// gc任务的状态，Success：已完成，Stopped：已停止，Running：清理中，Pending：排队中，Error：失败。
	JobStatus *ShowGcTaskResponseJobStatus `json:"job_status,omitempty"`

	// gc任务的创建时间。
	CreationTime *string `json:"creation_time,omitempty"`

	// gc任务的更新时间。
	UpdateTime     *string `json:"update_time,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowGcTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGcTaskResponse struct{}"
	}

	return strings.Join([]string{"ShowGcTaskResponse", string(data)}, " ")
}

type ShowGcTaskResponseJobKind struct {
	value string
}

type ShowGcTaskResponseJobKindEnum struct {
	MANUAL   ShowGcTaskResponseJobKind
	SCHEDULE ShowGcTaskResponseJobKind
}

func GetShowGcTaskResponseJobKindEnum() ShowGcTaskResponseJobKindEnum {
	return ShowGcTaskResponseJobKindEnum{
		MANUAL: ShowGcTaskResponseJobKind{
			value: "MANUAL",
		},
		SCHEDULE: ShowGcTaskResponseJobKind{
			value: "SCHEDULE",
		},
	}
}

func (c ShowGcTaskResponseJobKind) Value() string {
	return c.value
}

func (c ShowGcTaskResponseJobKind) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowGcTaskResponseJobKind) UnmarshalJSON(b []byte) error {
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

type ShowGcTaskResponseJobStatus struct {
	value string
}

type ShowGcTaskResponseJobStatusEnum struct {
	SUCCESS ShowGcTaskResponseJobStatus
	STOPPED ShowGcTaskResponseJobStatus
	RUNNING ShowGcTaskResponseJobStatus
	PENDING ShowGcTaskResponseJobStatus
	ERROR   ShowGcTaskResponseJobStatus
}

func GetShowGcTaskResponseJobStatusEnum() ShowGcTaskResponseJobStatusEnum {
	return ShowGcTaskResponseJobStatusEnum{
		SUCCESS: ShowGcTaskResponseJobStatus{
			value: "Success",
		},
		STOPPED: ShowGcTaskResponseJobStatus{
			value: "Stopped",
		},
		RUNNING: ShowGcTaskResponseJobStatus{
			value: "Running",
		},
		PENDING: ShowGcTaskResponseJobStatus{
			value: "Pending",
		},
		ERROR: ShowGcTaskResponseJobStatus{
			value: "Error",
		},
	}
}

func (c ShowGcTaskResponseJobStatus) Value() string {
	return c.value
}

func (c ShowGcTaskResponseJobStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowGcTaskResponseJobStatus) UnmarshalJSON(b []byte) error {
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
