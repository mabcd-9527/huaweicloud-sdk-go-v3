package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlarmSubRequest **参数解释**： 创建订阅告警请求体。 **取值范围**： 不涉及。
type AlarmSubRequest struct {

	// **参数解释**： 告警订阅名称。 **取值范围**： 不涉及。
	Name string `json:"name"`

	// **参数解释**： 是否开启订阅。 **取值范围**： 不涉及。
	Enable *int32 `json:"enable,omitempty"`

	// **参数解释**： 告警级别。 **取值范围**： 不涉及。
	AlarmLevel *string `json:"alarm_level,omitempty"`

	// **参数解释**： 消息主题地址。 **取值范围**： 不涉及。
	NotificationTarget string `json:"notification_target"`

	// **参数解释**： 消息主题名称。 **取值范围**： 不涉及。
	NotificationTargetName string `json:"notification_target_name"`

	// **参数解释**： 消息主题类型。 **取值范围**： - SMN：SMN类型
	NotificationTargetType string `json:"notification_target_type"`

	// **参数解释**： 时区。 **取值范围**： 不涉及。
	TimeZone string `json:"time_zone"`

	// **参数解释**： 集群ID。 **取值范围**： 不涉及。
	ClusterId *string `json:"cluster_id,omitempty"`

	// **参数解释**： 细粒度告警详细信息，规则为告警ID与告警级别用英文冒号分割，告警之间用逗号分割。当该字段不为空时，alarm_level字段需要设置为空 **取值范围**： 不涉及。
	AlarmDetail *string `json:"alarm_detail,omitempty"`
}

func (o AlarmSubRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmSubRequest struct{}"
	}

	return strings.Join([]string{"AlarmSubRequest", string(data)}, " ")
}
