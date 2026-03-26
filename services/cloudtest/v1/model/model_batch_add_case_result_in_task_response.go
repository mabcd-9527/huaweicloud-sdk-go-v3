package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAddCaseResultInTaskResponse Response Object
type BatchAddCaseResultInTaskResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchAddCaseResultInTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddCaseResultInTaskResponse struct{}"
	}

	return strings.Join([]string{"BatchAddCaseResultInTaskResponse", string(data)}, " ")
}
