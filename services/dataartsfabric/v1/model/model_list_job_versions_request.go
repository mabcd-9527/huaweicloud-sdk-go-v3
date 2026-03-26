package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListJobVersionsRequest Request Object
type ListJobVersionsRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 作业 ID
	JobId string `json:"job_id"`

	// 偏移量，表示从此偏移量开始查询， offset大于等于0，默认为0。
	Offset *int32 `json:"offset,omitempty"`

	// 指定每一页返回的最大条目数，取值范围[1,100]，默认为10。
	Limit *int32 `json:"limit,omitempty"`

	// version ID
	VersionId *string `json:"version_id,omitempty"`

	// Job version name，支持模糊查询
	Name *string `json:"name,omitempty"`
}

func (o ListJobVersionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobVersionsRequest struct{}"
	}

	return strings.Join([]string{"ListJobVersionsRequest", string(data)}, " ")
}
