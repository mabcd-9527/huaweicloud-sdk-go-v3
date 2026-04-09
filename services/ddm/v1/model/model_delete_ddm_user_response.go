package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDdmUserResponse Response Object
type DeleteDdmUserResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDdmUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDdmUserResponse struct{}"
	}

	return strings.Join([]string{"DeleteDdmUserResponse", string(data)}, " ")
}
