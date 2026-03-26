package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBackupConfigResponse Response Object
type ShowBackupConfigResponse struct {

	// 备份方式 - EBACKUP - PHYSICALBACKUP
	DefaultBackupMethod *string `json:"default_backup_method,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o ShowBackupConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBackupConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowBackupConfigResponse", string(data)}, " ")
}
