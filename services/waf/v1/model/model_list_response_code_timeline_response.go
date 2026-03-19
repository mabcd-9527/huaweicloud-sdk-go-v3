package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResponseCodeTimelineResponse Response Object
type ListResponseCodeTimelineResponse struct {

	// 安全统计的时间线，按时间顺序展示统计数据
	Body           *[]StatisticsTimelineItem `json:"body,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListResponseCodeTimelineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResponseCodeTimelineResponse struct{}"
	}

	return strings.Join([]string{"ListResponseCodeTimelineResponse", string(data)}, " ")
}
