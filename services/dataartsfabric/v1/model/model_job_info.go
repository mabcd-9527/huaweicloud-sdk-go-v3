package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobInfo Job详情信息
type JobInfo struct {

	// 可见性：  - PRIVATE：私有  - PUBLIC：公共  默认为PRIVATE
	Visibility *string `json:"visibility,omitempty"`

	// 作业Id
	Id *string `json:"id,omitempty"`

	// Job名称,只能包含中文、字母、数字、下划线、中划线、点、空格
	Name *string `json:"name,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 工作空间ID。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	Type *JobType `json:"type,omitempty"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	CurrentVersion *JobVersionInfo `json:"current_version,omitempty"`

	// Endpoint Id，32~36位的英文、数字、短横组合。
	CurrentEndpointId *string `json:"current_endpoint_id,omitempty"`

	CreateUser *User `json:"create_user,omitempty"`

	UpdateUser *User `json:"update_user,omitempty"`
}

func (o JobInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobInfo struct{}"
	}

	return strings.Join([]string{"JobInfo", string(data)}, " ")
}
