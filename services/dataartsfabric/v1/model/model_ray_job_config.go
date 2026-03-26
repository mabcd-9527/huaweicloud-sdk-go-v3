package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RayJobConfig Ray job类型的配置。
type RayJobConfig struct {

	// 运行作业主脚本
	Entrypoint string `json:"entrypoint"`

	// 用户提供的元数据
	Metadata *interface{} `json:"metadata,omitempty"`

	// 执行命令的CPU数量
	EntrypointNumCpus *float64 `json:"entrypoint_num_cpus,omitempty"`

	// 执行命令的GPU数量
	EntrypointNumGpus *float64 `json:"entrypoint_num_gpus,omitempty"`

	// 执行命令的内存大小
	EntrypointMemory *int32 `json:"entrypoint_memory,omitempty"`

	// 自定义资源数量
	EntrypointResources *interface{} `json:"entrypoint_resources,omitempty"`

	// 为作业指定的可选submission_id。
	SubmissionId *string `json:"submission_id,omitempty"`

	RuntimeEnv *RuntimeEnv `json:"runtime_env,omitempty"`
}

func (o RayJobConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RayJobConfig struct{}"
	}

	return strings.Join([]string{"RayJobConfig", string(data)}, " ")
}
