package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SwitchLoggerReplicaRequestBody struct {

	// **参数解释**:   切换日志副本的可用区。 **约束限制**:   不涉及。 **取值范围**:   不涉及。 **默认取值**:   不涉及。
	AvailabilityZone string `json:"availability_zone"`
}

func (o SwitchLoggerReplicaRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchLoggerReplicaRequestBody struct{}"
	}

	return strings.Join([]string{"SwitchLoggerReplicaRequestBody", string(data)}, " ")
}
