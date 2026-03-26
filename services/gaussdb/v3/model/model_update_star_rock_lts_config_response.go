package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateStarRockLtsConfigResponse Response Object
type UpdateStarRockLtsConfigResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UpdateStarRockLtsConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateStarRockLtsConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateStarRockLtsConfigResponse", string(data)}, " ")
}
