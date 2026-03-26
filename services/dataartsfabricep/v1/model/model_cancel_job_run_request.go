package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelJobRunRequest Request Object
type CancelJobRunRequest struct {

	// Workspace的ID
	WorkspaceId string `json:"workspace_id"`

	// 作业运行的ID
	RunId string `json:"run_id"`
}

func (o CancelJobRunRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelJobRunRequest struct{}"
	}

	return strings.Join([]string{"CancelJobRunRequest", string(data)}, " ")
}
