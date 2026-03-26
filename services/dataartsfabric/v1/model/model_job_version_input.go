package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobVersionInput 版本信息
type JobVersionInput struct {

	// 作业版本名称,只能包含中文、字母、数字、下划线、中划线、点、空格
	Name string `json:"name"`

	// cap白名单
	CapWhiteList *[]string `json:"cap_white_list,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	Config *JobConfig `json:"config"`

	// Framework列表信息
	DependenceFramework *[]Framework `json:"dependence_framework,omitempty"`
}

func (o JobVersionInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobVersionInput struct{}"
	}

	return strings.Join([]string{"JobVersionInput", string(data)}, " ")
}
