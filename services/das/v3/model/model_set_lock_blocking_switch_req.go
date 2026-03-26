package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SetLockBlockingSwitchReq struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 开关
	SwitchOn bool `json:"switch_on"`

	// 引擎
	EngineType string `json:"engine_type"`

	// 保存时长
	RetentionHours *int64 `json:"retention_hours,omitempty"`
}

func (o SetLockBlockingSwitchReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetLockBlockingSwitchReq struct{}"
	}

	return strings.Join([]string{"SetLockBlockingSwitchReq", string(data)}, " ")
}
