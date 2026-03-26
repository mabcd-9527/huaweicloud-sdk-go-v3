package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClearJobRunResponse Response Object
type ClearJobRunResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ClearJobRunResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClearJobRunResponse struct{}"
	}

	return strings.Join([]string{"ClearJobRunResponse", string(data)}, " ")
}
