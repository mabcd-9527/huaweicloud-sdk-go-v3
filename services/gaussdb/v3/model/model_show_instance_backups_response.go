package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceBackupsResponse Response Object
type ShowInstanceBackupsResponse struct {

	// **参数解释**： 指定实例全量备份信息列表。
	Backups *[]BackupV3 `json:"backups,omitempty"`

	// **参数解释**： 总备份数量。 **取值范围**： 不涉及。
	TotalCount     *string `json:"total_count,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowInstanceBackupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceBackupsResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceBackupsResponse", string(data)}, " ")
}
