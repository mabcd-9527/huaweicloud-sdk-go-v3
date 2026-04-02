package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGcTasksResponse Response Object
type ListGcTasksResponse struct {
	PageInfo *PageInfo `json:"page_info,omitempty"`

	// gc任务列表详情
	Tasks          *[]GcTask `json:"tasks,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListGcTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGcTasksResponse struct{}"
	}

	return strings.Join([]string{"ListGcTasksResponse", string(data)}, " ")
}
