package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobVersionInfo Job详情信息
type JobVersionInfo struct {

	// 作业版本名称,只能包含中文、字母、数字、下划线、中划线、点、空格
	Name string `json:"name"`

	// cap白名单
	CapWhiteList *[]string `json:"cap_white_list,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	Config *JobConfig `json:"config"`

	// Framework列表信息
	DependenceFramework *[]Framework `json:"dependence_framework,omitempty"`

	// 作业版本Id
	Id *string `json:"id,omitempty"`

	CreateUser *User `json:"create_user,omitempty"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`
}

func (o JobVersionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobVersionInfo struct{}"
	}

	return strings.Join([]string{"JobVersionInfo", string(data)}, " ")
}
