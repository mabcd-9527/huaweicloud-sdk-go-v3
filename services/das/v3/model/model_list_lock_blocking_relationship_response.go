package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLockBlockingRelationshipResponse Response Object
type ListLockBlockingRelationshipResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 锁阻塞明细列表
	DetailList     *[]ListLockBlockingDetailRespDetailList `json:"detail_list,omitempty"`
	HttpStatusCode int                                     `json:"-"`
}

func (o ListLockBlockingRelationshipResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLockBlockingRelationshipResponse struct{}"
	}

	return strings.Join([]string{"ListLockBlockingRelationshipResponse", string(data)}, " ")
}
