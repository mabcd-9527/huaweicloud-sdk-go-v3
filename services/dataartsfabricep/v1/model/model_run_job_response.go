package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunJobResponse Response Object
type RunJobResponse struct {

	// 作业运行Id
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RunJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunJobResponse struct{}"
	}

	return strings.Join([]string{"RunJobResponse", string(data)}, " ")
}
