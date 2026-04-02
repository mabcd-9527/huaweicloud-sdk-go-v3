package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifySecurityGroupResponse Response Object
type ModifySecurityGroupResponse struct {

	// **参数解释**: 修改实例安全组的任务ID。 **取值范围**: 不涉及。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ModifySecurityGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifySecurityGroupResponse struct{}"
	}

	return strings.Join([]string{"ModifySecurityGroupResponse", string(data)}, " ")
}
