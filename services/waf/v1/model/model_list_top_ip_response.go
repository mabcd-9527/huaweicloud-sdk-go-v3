package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTopIpResponse Response Object
type ListTopIpResponse struct {

	// CountItem的总数量
	Total *int32 `json:"total,omitempty"`

	// CountItem详细信息列表
	Items          *[]CountItem `json:"items,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListTopIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTopIpResponse struct{}"
	}

	return strings.Join([]string{"ListTopIpResponse", string(data)}, " ")
}
