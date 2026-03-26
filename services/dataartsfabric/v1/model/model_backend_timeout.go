package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BackendTimeout - **参数解释**：Ray Service独立Api设置的接口超时时间。 - **约束限制**：不涉及。 - **取值范围**：[0,600000]。 - **默认取值**：60000。
type BackendTimeout struct {
}

func (o BackendTimeout) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackendTimeout struct{}"
	}

	return strings.Join([]string{"BackendTimeout", string(data)}, " ")
}
