package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TopDomainsCountItem struct {

	// **参数解释：** 域名ID **取值范围：** 不涉及
	Key *string `json:"key,omitempty"`

	// **参数解释：** 数量 **取值范围：** 不涉及
	Num *int64 `json:"num,omitempty"`

	// **参数解释：** 网站名称，对应WAF控制台域名详情中的网站名称 **取值范围：** 不涉及
	WebTag *string `json:"web_tag,omitempty"`
}

func (o TopDomainsCountItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopDomainsCountItem struct{}"
	}

	return strings.Join([]string{"TopDomainsCountItem", string(data)}, " ")
}
