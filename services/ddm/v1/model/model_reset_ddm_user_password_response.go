package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetDdmUserPasswordResponse Response Object
type ResetDdmUserPasswordResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ResetDdmUserPasswordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetDdmUserPasswordResponse struct{}"
	}

	return strings.Join([]string{"ResetDdmUserPasswordResponse", string(data)}, " ")
}
