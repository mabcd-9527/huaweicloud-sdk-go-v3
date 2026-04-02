package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PremiumWafPoolResponse struct {

	// 实例组id
	Id *string `json:"id,omitempty"`

	// 实例组名称
	Name *string `json:"name,omitempty"`

	// 实例组所在region
	Region *string `json:"region,omitempty"`

	// 实例组类型
	Type *string `json:"type,omitempty"`

	// 实例组关联的vpc_id
	VpcId *string `json:"vpc_id,omitempty"`

	// 实例组描述
	Description *string `json:"description,omitempty"`

	// 实例组关联的防护域名
	Hosts *[]IdNameEntry `json:"hosts,omitempty"`

	// 实例组关联的引擎实例
	Instances *[]IdNameEntry `json:"instances,omitempty"`

	// 实例组关联的企业计划id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 实例组创建时间
	CreateTime *int64 `json:"create_time,omitempty"`
}

func (o PremiumWafPoolResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PremiumWafPoolResponse struct{}"
	}

	return strings.Join([]string{"PremiumWafPoolResponse", string(data)}, " ")
}
