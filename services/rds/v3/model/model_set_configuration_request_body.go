package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SetConfigurationRequestBody struct {

	// 备份方式 - EBACKUP - PHYSICALBACKUP
	DefaultBackupMethod string `json:"default_backup_method"`
}

func (o SetConfigurationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetConfigurationRequestBody struct{}"
	}

	return strings.Join([]string{"SetConfigurationRequestBody", string(data)}, " ")
}
