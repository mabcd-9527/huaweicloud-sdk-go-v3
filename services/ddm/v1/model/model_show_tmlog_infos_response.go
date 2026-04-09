package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTmlogInfosResponse Response Object
type ShowTmlogInfosResponse struct {

	// TMLOG信息列表。
	Tmlogs *[]TmlogInfo `json:"tmlogs,omitempty"`

	// 查询到TMLOG的节点总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowTmlogInfosResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTmlogInfosResponse struct{}"
	}

	return strings.Join([]string{"ShowTmlogInfosResponse", string(data)}, " ")
}
