package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClearJobRunRequest Request Object
type ClearJobRunRequest struct {

	// Workspace的ID
	WorkspaceId string `json:"workspace_id"`

	// Endpoint的ID
	EndpointId string `json:"endpoint_id"`
}

func (o ClearJobRunRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClearJobRunRequest struct{}"
	}

	return strings.Join([]string{"ClearJobRunRequest", string(data)}, " ")
}
