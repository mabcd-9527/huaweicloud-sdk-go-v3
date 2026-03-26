package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobName Job名称,只能包含中文、字母、数字、下划线、中划线、点、空格
type JobName struct {
}

func (o JobName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobName struct{}"
	}

	return strings.Join([]string{"JobName", string(data)}, " ")
}
