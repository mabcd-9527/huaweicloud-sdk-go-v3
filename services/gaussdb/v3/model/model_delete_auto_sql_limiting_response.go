package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAutoSqlLimitingResponse Response Object
type DeleteAutoSqlLimitingResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAutoSqlLimitingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAutoSqlLimitingResponse struct{}"
	}

	return strings.Join([]string{"DeleteAutoSqlLimitingResponse", string(data)}, " ")
}
