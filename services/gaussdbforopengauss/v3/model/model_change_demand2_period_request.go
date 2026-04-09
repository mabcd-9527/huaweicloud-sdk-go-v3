package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeDemand2PeriodRequest Request Object
type ChangeDemand2PeriodRequest struct {
	Body *ModifyDbPayTypeRequestBody `json:"body,omitempty"`
}

func (o ChangeDemand2PeriodRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeDemand2PeriodRequest struct{}"
	}

	return strings.Join([]string{"ChangeDemand2PeriodRequest", string(data)}, " ")
}
