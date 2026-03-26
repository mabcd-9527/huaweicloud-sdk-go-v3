package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListElbsResponse Response Object
type ListElbsResponse struct {

	// **参数解释**： 弹性负载均衡列表。 **取值范围**： 不涉及。
	Elbs *[]ClusterElbInfo `json:"elbs,omitempty"`

	// **参数解释**： 总条数。 **取值范围**： 不涉及。
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListElbsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListElbsResponse struct{}"
	}

	return strings.Join([]string{"ListElbsResponse", string(data)}, " ")
}
