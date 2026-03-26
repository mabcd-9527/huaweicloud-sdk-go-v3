package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobRef 引用定义
type JobRef struct {

	// 作业Id
	Id string `json:"id"`

	// 作业版本Id
	VersionId string `json:"version_id"`
}

func (o JobRef) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobRef struct{}"
	}

	return strings.Join([]string{"JobRef", string(data)}, " ")
}
