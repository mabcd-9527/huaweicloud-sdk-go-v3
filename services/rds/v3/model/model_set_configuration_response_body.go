package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SetConfigurationResponseBody struct {
}

func (o SetConfigurationResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetConfigurationResponseBody struct{}"
	}

	return strings.Join([]string{"SetConfigurationResponseBody", string(data)}, " ")
}
