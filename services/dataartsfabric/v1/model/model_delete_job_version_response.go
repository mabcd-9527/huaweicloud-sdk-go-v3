package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteJobVersionResponse Response Object
type DeleteJobVersionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteJobVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteJobVersionResponse struct{}"
	}

	return strings.Join([]string{"DeleteJobVersionResponse", string(data)}, " ")
}
