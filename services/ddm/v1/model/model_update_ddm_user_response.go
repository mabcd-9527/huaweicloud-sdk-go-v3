package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDdmUserResponse Response Object
type UpdateDdmUserResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateDdmUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDdmUserResponse struct{}"
	}

	return strings.Join([]string{"UpdateDdmUserResponse", string(data)}, " ")
}
