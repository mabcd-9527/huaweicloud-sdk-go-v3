package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAddCaseResultInTaskRequest Request Object
type BatchAddCaseResultInTaskRequest struct {

	// 项目id
	ProjectId string `json:"project_id"`

	// 版本URI
	VersionUri string `json:"version_uri"`

	Body *BatchAddTestCaseResultInTaskInfo `json:"body,omitempty"`
}

func (o BatchAddCaseResultInTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddCaseResultInTaskRequest struct{}"
	}

	return strings.Join([]string{"BatchAddCaseResultInTaskRequest", string(data)}, " ")
}
