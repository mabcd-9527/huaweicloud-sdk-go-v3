package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeDemand2PeriodResponse Response Object
type ChangeDemand2PeriodResponse struct {

	// **参数解释**: 订单ID的集合。
	OrderIds *[]string `json:"order_ids,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeDemand2PeriodResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeDemand2PeriodResponse struct{}"
	}

	return strings.Join([]string{"ChangeDemand2PeriodResponse", string(data)}, " ")
}
