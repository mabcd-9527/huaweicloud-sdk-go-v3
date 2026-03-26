package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobRunBriefInfo JobRun简要信息
type JobRunBriefInfo struct {
	Job *JobRef `json:"job"`

	// 作业运行的名称，只能包含中文、字母、数字、下划线、中划线、点、空格
	Name string `json:"name"`

	// endpoint空间id
	EndpointId string `json:"endpoint_id"`

	// 一种资源ID，32~36位的英文、数字、短横组合
	Id string `json:"id"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time"`

	// 更新时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	Status *StatusEnum `json:"status"`

	// 运行时间
	Duration *int64 `json:"duration,omitempty"`

	Type *JobType `json:"type,omitempty"`

	CreateUser *User `json:"create_user,omitempty"`
}

func (o JobRunBriefInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobRunBriefInfo struct{}"
	}

	return strings.Join([]string{"JobRunBriefInfo", string(data)}, " ")
}
