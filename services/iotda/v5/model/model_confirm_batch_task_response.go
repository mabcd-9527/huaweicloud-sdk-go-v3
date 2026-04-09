package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfirmBatchTaskResponse Response Object
type ConfirmBatchTaskResponse struct {

	// 批量操作目标结果集合
	Targets        *[]BatchTargetResult `json:"targets,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ConfirmBatchTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmBatchTaskResponse struct{}"
	}

	return strings.Join([]string{"ConfirmBatchTaskResponse", string(data)}, " ")
}
