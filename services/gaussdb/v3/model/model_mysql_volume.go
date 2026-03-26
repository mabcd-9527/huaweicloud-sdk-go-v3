package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MysqlVolume struct {

	// **参数解释**：  磁盘大小，单位GB。  **约束限制**：  必须为10的整数倍。创建按需实例时不可选。  **取值范围**：  10-128000。  **默认取值**：  10。
	Size *string `json:"size,omitempty"`

	// **参数解释**：  磁盘存储类型。  **约束限制**：  不涉及。  **取值范围**：  - DL6：DL6存储类型。 - DL5：DL5存储类型。  **默认取值**：  DL6。
	Type *string `json:"type,omitempty"`
}

func (o MysqlVolume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MysqlVolume struct{}"
	}

	return strings.Join([]string{"MysqlVolume", string(data)}, " ")
}
