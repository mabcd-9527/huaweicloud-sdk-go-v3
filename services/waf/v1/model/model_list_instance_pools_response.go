package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstancePoolsResponse Response Object
type ListInstancePoolsResponse struct {

	// 实例组总数
	Total *int64 `json:"total,omitempty"`

	// 实例组列表
	Items          *[]PremiumWafPoolResponse `json:"items,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListInstancePoolsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancePoolsResponse struct{}"
	}

	return strings.Join([]string{"ListInstancePoolsResponse", string(data)}, " ")
}
