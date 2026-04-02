package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PremiumWafInstanceStatusResponse struct {

	// 实例id
	Id *string `json:"id,omitempty"`

	// 实例名称
	InstanceName *string `json:"instance_name,omitempty"`

	// **参数解释：** 实例计费状态 **约束限制：** 不涉及 **取值范围：** - -1 退订 - 0 正常 - 1 冻结 - 2 终止 - 3 受限 **默认取值：** 不涉及
	Status *int32 `json:"status,omitempty"`

	// **参数解释：** 实例运行状态 **约束限制：** 不涉及 **取值范围：** - 0 创建中 - 1 运行中 - 2 删除中 - 3 已删除 - 4 创建失败 - 5 已冻结 - 6 异常 - 7 更新中 - 8 更新失败 **默认取值：** 不涉及
	RunStatus *int32 `json:"run_status,omitempty"`

	// **参数解释：** 实例接入状态 **约束限制：** 不涉及 **取值范围：** - 0 未接入 - 1 已接入 - 2 DNS解析异常 **默认取值：** 不涉及
	AccessStatus *int32 `json:"access_status,omitempty"`
}

func (o PremiumWafInstanceStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PremiumWafInstanceStatusResponse struct{}"
	}

	return strings.Join([]string{"PremiumWafInstanceStatusResponse", string(data)}, " ")
}
