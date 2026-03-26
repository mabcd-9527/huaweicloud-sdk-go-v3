package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteStarRockLtsConfigResponse Response Object
type DeleteStarRockLtsConfigResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o DeleteStarRockLtsConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteStarRockLtsConfigResponse struct{}"
	}

	return strings.Join([]string{"DeleteStarRockLtsConfigResponse", string(data)}, " ")
}
