package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteJobRunRequest Request Object
type DeleteJobRunRequest struct {

	// Workspace的ID
	WorkspaceId string `json:"workspace_id"`

	// 作业运行的ID
	RunId string `json:"run_id"`
}

func (o DeleteJobRunRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteJobRunRequest struct{}"
	}

	return strings.Join([]string{"DeleteJobRunRequest", string(data)}, " ")
}
