package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RuntimeEnv 作业的运行时环境配置，可选参数有：working_dir：代码将在其中运行的工作目录。必须是远程URI，如s3或git路径。py_modules：将与运行时环境一起安装的Python模块。这些必须是远程URI。pip：要安装的pip软件包列表。conda：conda YAML配置或本地conda env的名称（例如，“pytorch_p36”），env_vars：要设置的环境变量。
type RuntimeEnv struct {

	// 代码将在其中运行的工作目录。必须是远程URI，如s3或git路径。
	WorkingDir *string `json:"working_dir,omitempty"`

	// 将与运行时环境一起安装的Python模块。这些必须是远程URI。
	PyModules *[]string `json:"py_modules,omitempty"`

	// 要安装的pip软件包列表。
	Pip *[]string `json:"pip,omitempty"`

	// 【联合【指令【str，任何】，str】：conda YAML配置或本地conda env的名称（例如，“pytorch_p36”），
	Conda *interface{} `json:"conda,omitempty"`

	// 要设置的环境变量。
	EnvVars map[string]string `json:"env_vars,omitempty"`

	// 容器配置
	Container *interface{} `json:"container,omitempty"`

	// hook配置信息
	WorkerProcessSetupHook *string `json:"worker_process_setup_hook,omitempty"`

	// nsight配置
	Nsight *string `json:"nsight,omitempty"`

	// 运行环境的配置
	Config *string `json:"config,omitempty"`
}

func (o RuntimeEnv) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuntimeEnv struct{}"
	}

	return strings.Join([]string{"RuntimeEnv", string(data)}, " ")
}
