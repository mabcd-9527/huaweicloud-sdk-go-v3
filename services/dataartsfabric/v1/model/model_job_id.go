package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobId 作业Id
type JobId struct {
}

func (o JobId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobId struct{}"
	}

	return strings.Join([]string{"JobId", string(data)}, " ")
}
