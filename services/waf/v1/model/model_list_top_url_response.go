package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTopUrlResponse Response Object
type ListTopUrlResponse struct {

	// UrlCountItem的总数量
	Total *int32 `json:"total,omitempty"`

	// UrlCountItem详细信息列表
	Items          *[]UrlCountItem `json:"items,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListTopUrlResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTopUrlResponse struct{}"
	}

	return strings.Join([]string{"ListTopUrlResponse", string(data)}, " ")
}
