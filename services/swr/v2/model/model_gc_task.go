package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type GcTask struct {

	// gc任务的ID。
	Id *int32 `json:"id,omitempty"`

	// gc任务的名称。
	JobName *string `json:"job_name,omitempty"`

	// 任务类型，MANUAL：手动执行，SCHEDULE：定时执行。
	JobKind *GcTaskJobKind `json:"job_kind,omitempty"`

	JobParameters *JobParameters `json:"job_parameters,omitempty"`

	// gc任务的状态，Success：已完成，Stopped：已停止，Running：清理中，Pending：排队中，Error：失败。
	JobStatus *GcTaskJobStatus `json:"job_status,omitempty"`

	// gc任务的创建时间。
	CreationTime *string `json:"creation_time,omitempty"`

	// gc任务的更新时间。
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o GcTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GcTask struct{}"
	}

	return strings.Join([]string{"GcTask", string(data)}, " ")
}

type GcTaskJobKind struct {
	value string
}

type GcTaskJobKindEnum struct {
	MANUAL   GcTaskJobKind
	SCHEDULE GcTaskJobKind
}

func GetGcTaskJobKindEnum() GcTaskJobKindEnum {
	return GcTaskJobKindEnum{
		MANUAL: GcTaskJobKind{
			value: "MANUAL",
		},
		SCHEDULE: GcTaskJobKind{
			value: "SCHEDULE",
		},
	}
}

func (c GcTaskJobKind) Value() string {
	return c.value
}

func (c GcTaskJobKind) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GcTaskJobKind) UnmarshalJSON(b []byte) error {
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

type GcTaskJobStatus struct {
	value string
}

type GcTaskJobStatusEnum struct {
	SUCCESS GcTaskJobStatus
	STOPPED GcTaskJobStatus
	RUNNING GcTaskJobStatus
	PENDING GcTaskJobStatus
	ERROR   GcTaskJobStatus
}

func GetGcTaskJobStatusEnum() GcTaskJobStatusEnum {
	return GcTaskJobStatusEnum{
		SUCCESS: GcTaskJobStatus{
			value: "Success",
		},
		STOPPED: GcTaskJobStatus{
			value: "Stopped",
		},
		RUNNING: GcTaskJobStatus{
			value: "Running",
		},
		PENDING: GcTaskJobStatus{
			value: "Pending",
		},
		ERROR: GcTaskJobStatus{
			value: "Error",
		},
	}
}

func (c GcTaskJobStatus) Value() string {
	return c.value
}

func (c GcTaskJobStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GcTaskJobStatus) UnmarshalJSON(b []byte) error {
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
