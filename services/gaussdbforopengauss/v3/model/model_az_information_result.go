package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AzInformationResult struct {

	// **参数解释**: 所属AZ代码。 **取值范围**: 不涉及。
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// **参数解释**: 所属AZ名称。 **取值范围**: 不涉及。
	Description *string `json:"description,omitempty"`
}

func (o AzInformationResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AzInformationResult struct{}"
	}

	return strings.Join([]string{"AzInformationResult", string(data)}, " ")
}
