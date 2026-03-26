package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteJobVersionRequest Request Object
type DeleteJobVersionRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 作业 ID
	JobId string `json:"job_id"`

	// 作业版本ID
	VersionId string `json:"version_id"`
}

func (o DeleteJobVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteJobVersionRequest struct{}"
	}

	return strings.Join([]string{"DeleteJobVersionRequest", string(data)}, " ")
}
