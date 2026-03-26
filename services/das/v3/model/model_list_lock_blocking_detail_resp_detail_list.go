package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListLockBlockingDetailRespDetailList struct {

	// 唯一ID
	Id *string `json:"id,omitempty"`

	// 批次ID
	UniqueId *string `json:"unique_id,omitempty"`

	// 会话ID
	Spid *int64 `json:"spid,omitempty"`

	// 阻塞会话ID
	BlockedBySpid *int64 `json:"blocked_by_spid,omitempty"`

	// 采集时间
	CollectTime *int64 `json:"collect_time,omitempty"`

	// SQL开始时间
	StartTime *int64 `json:"start_time,omitempty"`

	// 执行耗时（ms）
	ElapsedTime *int64 `json:"elapsed_time,omitempty"`

	// CPU时间（ms）
	Cpu *int64 `json:"cpu,omitempty"`

	// 阻塞时长（ms）
	WaitTime *int64 `json:"wait_time,omitempty"`

	// 等待类型
	WaitType *string `json:"wait_type,omitempty"`

	// SQL
	Sql *string `json:"sql,omitempty"`

	// Query Hash
	QueryHash *string `json:"query_hash,omitempty"`

	// 数据库名
	DbName *string `json:"db_name,omitempty"`

	// 客户端HostName
	HostName *string `json:"host_name,omitempty"`

	// 客户端名称
	ClientAppName *string `json:"client_app_name,omitempty"`

	// 当前用户
	LoginId *string `json:"login_id,omitempty"`

	// SQL命令类型
	Cmd *string `json:"cmd,omitempty"`

	// SQL消耗的I/O
	PhysicalIo *int64 `json:"physical_io,omitempty"`
}

func (o ListLockBlockingDetailRespDetailList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLockBlockingDetailRespDetailList struct{}"
	}

	return strings.Join([]string{"ListLockBlockingDetailRespDetailList", string(data)}, " ")
}
