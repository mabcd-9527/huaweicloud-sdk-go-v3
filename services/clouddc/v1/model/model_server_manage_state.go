package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServerManageState **参数解释**： 服务器的生命周期状态及对应的数量 **约束限制**： 格式为key: value state(生命周期状态): count(当前生命周期状态的服务器数量)
type ServerManageState struct {
}

func (o ServerManageState) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerManageState struct{}"
	}

	return strings.Join([]string{"ServerManageState", string(data)}, " ")
}
