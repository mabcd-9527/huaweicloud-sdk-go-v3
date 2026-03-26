package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SparkJobConfig struct {

	// Job主文件，obs上的python文件或jar包
	MainFilePath *string `json:"main_file_path,omitempty"`

	// Java job需指定主类
	MainClass *string `json:"main_class,omitempty"`

	// Python主文件的运行参数或Java主类的运行参数
	MainArguments *[]string `json:"main_arguments,omitempty"`

	// 自定义的spark配置值
	SparkConf *string `json:"spark_conf,omitempty"`

	// driver资源规格
	DriverResourceSpec string `json:"driver_resource_spec"`

	// executor资源规格
	ExecutorResourceSpec string `json:"executor_resource_spec"`

	// executor数量
	NumExecutors *int32 `json:"num_executors,omitempty"`

	// 是否开启dynamic alocation
	EnableDynamicAllocation *bool `json:"enable_dynamic_allocation,omitempty"`

	// executor最少个数
	MinNumExecutors *int32 `json:"min_num_executors,omitempty"`

	// executor最多个数
	MaxNumExecutors *int32 `json:"max_num_executors,omitempty"`

	// 用户日志obs存储位置
	UserLogPath *string `json:"user_log_path,omitempty"`

	// Metastore信息，LakeFormation服务的实例Id，即MetaStoreId。
	MetastoreId *string `json:"metastore_id,omitempty"`

	// metastore catalog id，即lakeformation catalog id
	MetastoreCatalogId *string `json:"metastore_catalog_id,omitempty"`
}

func (o SparkJobConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SparkJobConfig struct{}"
	}

	return strings.Join([]string{"SparkJobConfig", string(data)}, " ")
}
