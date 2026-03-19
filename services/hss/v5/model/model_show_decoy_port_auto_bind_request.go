package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDecoyPortAutoBindRequest Request Object
type ShowDecoyPortAutoBindRequest struct {

	// **参数解释**: 企业项目ID，（该字段已废弃） **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ShowDecoyPortAutoBindRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDecoyPortAutoBindRequest struct{}"
	}

	return strings.Join([]string{"ShowDecoyPortAutoBindRequest", string(data)}, " ")
}
