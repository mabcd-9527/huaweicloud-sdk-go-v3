package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVulnRulesResponse Response Object
type ListVulnRulesResponse struct {
	Body           *[]VulnDto `json:"body,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListVulnRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVulnRulesResponse struct{}"
	}

	return strings.Join([]string{"ListVulnRulesResponse", string(data)}, " ")
}
