package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeBackupConfigResponse Response Object
type ChangeBackupConfigResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeBackupConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeBackupConfigResponse struct{}"
	}

	return strings.Join([]string{"ChangeBackupConfigResponse", string(data)}, " ")
}
