package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateYmlsReqEdit **参数解释**： 配置文件信息。 **约束限制**： 不涉及
type UpdateYmlsReqEdit struct {
	Modify *UpdateYmlsReqEditModify `json:"modify,omitempty"`

	Delete *UpdateYmlsReqEditModify `json:"delete,omitempty"`

	Reset *UpdateYmlsReqEditModify `json:"reset,omitempty"`
}

func (o UpdateYmlsReqEdit) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateYmlsReqEdit struct{}"
	}

	return strings.Join([]string{"UpdateYmlsReqEdit", string(data)}, " ")
}
