package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobConfig Job配置
type JobConfig struct {
	RayJobConfig *RayJobConfig `json:"ray_job_config,omitempty"`

	PythonJobConfig *PythonJobConfig `json:"python_job_config,omitempty"`

	SparkJobConfig *SparkJobConfig `json:"spark_job_config,omitempty"`

	// 超时时间，即最长运行时间,超过该事件则终止。   可选，默认值为0，值0表示没有超时。
	TimeoutSeconds *int32 `json:"timeout_seconds,omitempty"`

	// 运行重试次数
	MaxRetries *int32 `json:"max_retries,omitempty"`

	// 作业是否排队。 true: 排队 false：不排队
	Queue *bool `json:"queue,omitempty"`

	// 作业允许的最大并发运行数可选。 Default: 1
	MaxConcurrentRuns *int32 `json:"max_concurrent_runs,omitempty"`
}

func (o JobConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobConfig struct{}"
	}

	return strings.Join([]string{"JobConfig", string(data)}, " ")
}
