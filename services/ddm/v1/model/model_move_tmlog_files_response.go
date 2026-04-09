package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MoveTmlogFilesResponse Response Object
type MoveTmlogFilesResponse struct {

	// 任务流ID
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o MoveTmlogFilesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MoveTmlogFilesResponse struct{}"
	}

	return strings.Join([]string{"MoveTmlogFilesResponse", string(data)}, " ")
}
