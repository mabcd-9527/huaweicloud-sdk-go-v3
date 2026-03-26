package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AgencyConfig - **参数解释**：Ray Service委托配置信息。 - **约束限制**：不涉及。 - **取值范围**：不涉及。 - **默认取值**：不涉及。
type AgencyConfig struct {

	// - **参数解释**：是否开启Ray Service委托授权功能。 - **约束限制**：不涉及。 - **取值范围**：开启true，关闭false。 - **默认取值**：false。
	EnableAgency *bool `json:"enable_agency,omitempty"`

	// - **参数解释**：授予Ray Service的委托名称。 - **约束限制**：不超过64位。 - **取值范围**：不涉及。 - **默认取值**：null。
	AgencyName *string `json:"agency_name,omitempty"`
}

func (o AgencyConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgencyConfig struct{}"
	}

	return strings.Join([]string{"AgencyConfig", string(data)}, " ")
}
