package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunId 作业运行Id
type RunId struct {
}

func (o RunId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunId struct{}"
	}

	return strings.Join([]string{"RunId", string(data)}, " ")
}
