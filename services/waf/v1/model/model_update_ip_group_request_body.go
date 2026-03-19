package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateIpGroupRequestBody 地址组名称
type UpdateIpGroupRequestBody struct {

	// 地址组名称
	Name *string `json:"name,omitempty"`

	// **参数解释：** 以逗号分隔的ip或ip段 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Ips string `json:"ips"`

	// **参数解释：** ip或ip段的备注 **约束限制：** key必须是ips中包含的单个ip或ip段 **取值范围：** value必须匹配正则：[^<>]{0,64} **默认取值：** 不涉及
	IpRemarks map[string]string `json:"ip_remarks,omitempty"`

	// 地址组描述
	Description *string `json:"description,omitempty"`
}

func (o UpdateIpGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIpGroupRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateIpGroupRequestBody", string(data)}, " ")
}
