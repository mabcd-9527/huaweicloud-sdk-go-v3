package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateJobRequest Request Object
type UpdateJobRequest struct {

	// 作业 ID
	JobId string `json:"job_id"`

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	Body *UpdateJobInput `json:"body,omitempty"`
}

func (o UpdateJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateJobRequest struct{}"
	}

	return strings.Join([]string{"UpdateJobRequest", string(data)}, " ")
}
