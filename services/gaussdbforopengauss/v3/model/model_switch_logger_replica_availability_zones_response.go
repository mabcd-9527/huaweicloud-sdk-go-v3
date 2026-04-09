package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SwitchLoggerReplicaAvailabilityZonesResponse Response Object
type SwitchLoggerReplicaAvailabilityZonesResponse struct {

	// **参数解释**: 选择日志节点AZ列表。
	AvailabilityZoneInfos *[]AzInformationResult `json:"availability_zone_infos,omitempty"`
	HttpStatusCode        int                    `json:"-"`
}

func (o SwitchLoggerReplicaAvailabilityZonesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchLoggerReplicaAvailabilityZonesResponse struct{}"
	}

	return strings.Join([]string{"SwitchLoggerReplicaAvailabilityZonesResponse", string(data)}, " ")
}
