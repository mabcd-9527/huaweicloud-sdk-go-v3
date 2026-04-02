package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VulnDto 漏洞信息
type VulnDto struct {

	// 漏洞编号
	VulnerabilityId *string `json:"vulnerability_id,omitempty"`

	// 漏洞描述
	Description *string `json:"description,omitempty"`

	// 可防护的规则ID列表
	RuleIds *[]string `json:"rule_ids,omitempty"`

	// 漏洞情报创建时间
	CreateTime *int64 `json:"create_time,omitempty"`
}

func (o VulnDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VulnDto struct{}"
	}

	return strings.Join([]string{"VulnDto", string(data)}, " ")
}
