package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelJobRunResponse Response Object
type CancelJobRunResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CancelJobRunResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelJobRunResponse struct{}"
	}

	return strings.Join([]string{"CancelJobRunResponse", string(data)}, " ")
}
