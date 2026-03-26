package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListJobRunsResponse Response Object
type ListJobRunsResponse struct {

	// 符合条件的job Version总数
	Total *int32 `json:"total,omitempty"`

	// 符合条件的运行作业简要信息列表
	JobRuns *[]JobRunBriefInfo `json:"job_runs,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListJobRunsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobRunsResponse struct{}"
	}

	return strings.Join([]string{"ListJobRunsResponse", string(data)}, " ")
}
