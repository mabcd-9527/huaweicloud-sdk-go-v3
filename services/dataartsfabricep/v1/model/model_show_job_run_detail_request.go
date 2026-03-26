package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowJobRunDetailRequest Request Object
type ShowJobRunDetailRequest struct {

	// Workspace的ID
	WorkspaceId string `json:"workspace_id"`

	// 作业运行的ID
	RunId string `json:"run_id"`
}

func (o ShowJobRunDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobRunDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowJobRunDetailRequest", string(data)}, " ")
}
