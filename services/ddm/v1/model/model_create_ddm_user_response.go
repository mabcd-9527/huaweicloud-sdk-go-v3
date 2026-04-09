package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDdmUserResponse Response Object
type CreateDdmUserResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateDdmUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDdmUserResponse struct{}"
	}

	return strings.Join([]string{"CreateDdmUserResponse", string(data)}, " ")
}
