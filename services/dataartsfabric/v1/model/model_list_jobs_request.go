package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListJobsRequest Request Object
type ListJobsRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 偏移量，表示从此偏移量开始查询， offset大于等于0，默认为0。
	Offset *int32 `json:"offset,omitempty"`

	// 指定每一页返回的最大条目数，取值范围[1,100]，默认为10。
	Limit *int32 `json:"limit,omitempty"`

	// 通过名字搜索作业
	Name *string `json:"name,omitempty"`

	// 通过作业id检索
	Id *string `json:"id,omitempty"`

	// 通过类型检索
	Type *string `json:"type,omitempty"`

	// 通过id检索Endpoint的参数
	EndpointId *string `json:"endpoint_id,omitempty"`

	// 指定排序字段，可选值为： CREATE_TIME：创建时间，默认值 UPDATE_TIME：更新时间 NAME: 服务名称
	SortBy *string `json:"sort_by,omitempty"`

	// 排序方式，可选值如下： ASC : 递增排序 DESC: 递减排序，默认值
	OrderBy *string `json:"order_by,omitempty"`

	// 是否由我创建的作业。 默认为null，即返回租户下所有子用户的作业； True: 只返回当前子用户的作业； False: 返回租户下所有子用户的作业；
	CreateByMe *bool `json:"create_by_me,omitempty"`
}

func (o ListJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobsRequest struct{}"
	}

	return strings.Join([]string{"ListJobsRequest", string(data)}, " ")
}
