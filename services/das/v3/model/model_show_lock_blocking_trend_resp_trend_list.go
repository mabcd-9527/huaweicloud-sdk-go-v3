package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowLockBlockingTrendRespTrendList struct {

	// 采集时间
	CollectTime *int64 `json:"collect_time,omitempty"`

	// 锁阻塞总数
	TotalLockBlockingCount *int64 `json:"total_lock_blocking_count,omitempty"`

	// ASYNC_IO_COMPLETION数量
	AsyncIoCompletionCount *int64 `json:"async_io_completion_count,omitempty"`

	// ASYNC_NETWORK_IO数量
	AsyncNetworkIoCount *int64 `json:"async_network_io_count,omitempty"`

	// CXCONSUMER数量
	CxconsumerCount *int64 `json:"cxconsumer_count,omitempty"`

	// CXPACKET数量
	CxpacketCount *int64 `json:"cxpacket_count,omitempty"`

	// DTC数量
	DtcCount *int64 `json:"dtc_count,omitempty"`

	// LCK_M_BU数量
	LckMBuCount *int64 `json:"lck_m_bu_count,omitempty"`

	// LCK_M_IS数量
	LckMIsCount *int64 `json:"lck_m_is_count,omitempty"`

	// LCK_M_IU数量
	LckMIuCount *int64 `json:"lck_m_iu_count,omitempty"`

	// LCK_M_IX数量
	LckMIxCount *int64 `json:"lck_m_ix_count,omitempty"`

	// 其他锁阻塞数量
	OtherCount *int64 `json:"other_count,omitempty"`
}

func (o ShowLockBlockingTrendRespTrendList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLockBlockingTrendRespTrendList struct{}"
	}

	return strings.Join([]string{"ShowLockBlockingTrendRespTrendList", string(data)}, " ")
}
