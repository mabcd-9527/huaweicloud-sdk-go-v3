package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLockBlockingTrendResponse Response Object
type ShowLockBlockingTrendResponse struct {

	// 时间间隔（ms）
	IntervalMillis *int64 `json:"interval_millis,omitempty"`

	// 锁阻塞趋势列表
	TrendList      *[]ShowLockBlockingTrendRespTrendList `json:"trend_list,omitempty"`
	HttpStatusCode int                                   `json:"-"`
}

func (o ShowLockBlockingTrendResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLockBlockingTrendResponse struct{}"
	}

	return strings.Join([]string{"ShowLockBlockingTrendResponse", string(data)}, " ")
}
