package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// JobType 类别，可选值如下： RAY_JOB: Ray Job,  PYTHON_JOB: Python Job,  SPARK_JOB: Spark Job
type JobType struct {
	value string
}

type JobTypeEnum struct {
	RAY_JOB    JobType
	PYTHON_JOB JobType
	SPARK_JOB  JobType
}

func GetJobTypeEnum() JobTypeEnum {
	return JobTypeEnum{
		RAY_JOB: JobType{
			value: "RAY_JOB",
		},
		PYTHON_JOB: JobType{
			value: "PYTHON_JOB",
		},
		SPARK_JOB: JobType{
			value: "SPARK_JOB",
		},
	}
}

func (c JobType) Value() string {
	return c.value
}

func (c JobType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobType) UnmarshalJSON(b []byte) error {
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
