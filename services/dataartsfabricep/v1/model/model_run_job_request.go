package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunJobRequest Request Object
type RunJobRequest struct {

	// Workspace的ID
	WorkspaceId string `json:"workspace_id"`

	Body *RunJobInput `json:"body,omitempty"`
}

func (o RunJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunJobRequest struct{}"
	}

	return strings.Join([]string{"RunJobRequest", string(data)}, " ")
}
