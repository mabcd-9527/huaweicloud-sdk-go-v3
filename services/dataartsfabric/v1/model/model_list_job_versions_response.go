package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListJobVersionsResponse Response Object
type ListJobVersionsResponse struct {

	// 符合条件的job Version总数
	Total *int32 `json:"total,omitempty"`

	// 列表信息
	Versions *[]JobVersionInfo `json:"versions,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListJobVersionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobVersionsResponse struct{}"
	}

	return strings.Join([]string{"ListJobVersionsResponse", string(data)}, " ")
}
