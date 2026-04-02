package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceStatusResponse Response Object
type ListInstanceStatusResponse struct {

	// 是否已购买独享引擎
	Purchased *bool `json:"purchased,omitempty"`

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 引擎实例状态信息列表
	Items          *[]PremiumWafInstanceStatusResponse `json:"items,omitempty"`
	HttpStatusCode int                                 `json:"-"`
}

func (o ListInstanceStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceStatusResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceStatusResponse", string(data)}, " ")
}
