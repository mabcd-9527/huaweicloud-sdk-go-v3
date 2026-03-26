package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobVersionId 作业版本Id
type JobVersionId struct {
}

func (o JobVersionId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobVersionId struct{}"
	}

	return strings.Join([]string{"JobVersionId", string(data)}, " ")
}
