package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOverviewsHostDomainsResponse Response Object
type ListOverviewsHostDomainsResponse struct {

	// 子域名数量
	Total *int32 `json:"total,omitempty"`

	// 子域名列表
	Items          *[]string `json:"items,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListOverviewsHostDomainsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOverviewsHostDomainsResponse struct{}"
	}

	return strings.Join([]string{"ListOverviewsHostDomainsResponse", string(data)}, " ")
}
