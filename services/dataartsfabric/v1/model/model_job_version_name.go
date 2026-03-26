package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobVersionName 作业版本名称,只能包含中文、字母、数字、下划线、中划线、点、空格
type JobVersionName struct {
}

func (o JobVersionName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobVersionName struct{}"
	}

	return strings.Join([]string{"JobVersionName", string(data)}, " ")
}
