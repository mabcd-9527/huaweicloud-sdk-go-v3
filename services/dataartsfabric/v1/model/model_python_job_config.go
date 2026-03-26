package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PythonJobConfig Python Job Config
type PythonJobConfig struct {

	// Job主文件，python文件，拼接在working_dir后可以是相对路径
	MainFileName *string `json:"main_file_name,omitempty"`

	// Python主文件的运行参数
	MainArguments *[]string `json:"main_arguments,omitempty"`

	// 依赖的python包列表
	PythonPackages *[]string `json:"python_packages,omitempty"`

	// 代码目录，obs目录
	WorkingDir *string `json:"working_dir,omitempty"`

	// 资源规格
	ResourceSpec string `json:"resource_spec"`

	// 用户日志obs存储位置
	UserLogPath *string `json:"user_log_path,omitempty"`
}

func (o PythonJobConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PythonJobConfig struct{}"
	}

	return strings.Join([]string{"PythonJobConfig", string(data)}, " ")
}
