package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModifyDbPayTypeRequestBody struct {

	// **参数解释**: 需要转成包周期计费模式的实例ID列表。 **约束限制**: 不涉及。
	EntityIds []string `json:"entity_ids"`

	ChargeInfo *PeriodChargeInfoOption `json:"charge_info"`
}

func (o ModifyDbPayTypeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyDbPayTypeRequestBody struct{}"
	}

	return strings.Join([]string{"ModifyDbPayTypeRequestBody", string(data)}, " ")
}
