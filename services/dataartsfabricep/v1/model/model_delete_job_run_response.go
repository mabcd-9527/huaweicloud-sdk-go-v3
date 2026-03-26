package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteJobRunResponse Response Object
type DeleteJobRunResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteJobRunResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteJobRunResponse struct{}"
	}

	return strings.Join([]string{"DeleteJobRunResponse", string(data)}, " ")
}
