package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModifySecurityGroupRequestBody struct {

	// **参数解释**: 目标安全组ID。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	SecurityGroupId string `json:"security_group_id"`
}

func (o ModifySecurityGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifySecurityGroupRequestBody struct{}"
	}

	return strings.Join([]string{"ModifySecurityGroupRequestBody", string(data)}, " ")
}
