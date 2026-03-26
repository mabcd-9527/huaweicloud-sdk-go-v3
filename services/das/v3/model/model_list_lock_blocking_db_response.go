package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLockBlockingDbResponse Response Object
type ListLockBlockingDbResponse struct {

	// 数据库名列表
	DbList         *[]string `json:"db_list,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListLockBlockingDbResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLockBlockingDbResponse struct{}"
	}

	return strings.Join([]string{"ListLockBlockingDbResponse", string(data)}, " ")
}
