package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobRunName 作业运行的名称，只能包含中文、字母、数字、下划线、中划线、点、空格
type JobRunName struct {
}

func (o JobRunName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobRunName struct{}"
	}

	return strings.Join([]string{"JobRunName", string(data)}, " ")
}
