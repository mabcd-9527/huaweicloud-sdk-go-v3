package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OidcConfig - **参数解释**：Oidc配置信息。 - **约束限制**：不涉及。 - **取值范围**：不涉及。 - **默认取值**：不涉及。
type OidcConfig struct {

	// - **参数解释**：是否开启oidc功能。 - **约束限制**：不涉及。 - **取值范围**：开启true，关闭false。 - **默认取值**：false。
	EnableOidc *bool `json:"enable_oidc,omitempty"`

	// - **参数解释**：客户端ID。 - **约束限制**：不涉及。 - **取值范围**：长度为[1,64]的字母、数字、下划线(_)、中划线(-)、点(.)的组合。 - **默认取值**：null。
	ClientId *string `json:"client_id,omitempty"`

	// - **参数解释**：RayService命名空间。 - **约束限制**：不涉及。 - **取值范围**：长度为16位，均为ray-svc-xxx格式。 - **默认取值**：null。
	Namespace *string `json:"namespace,omitempty"`
}

func (o OidcConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OidcConfig struct{}"
	}

	return strings.Join([]string{"OidcConfig", string(data)}, " ")
}
