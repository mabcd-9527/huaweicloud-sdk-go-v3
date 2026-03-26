package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateJobResponse Response Object
type CreateJobResponse struct {

	// 作业Id
	Id *string `json:"id,omitempty"`

	// 作业版本Id
	VersionId *string `json:"version_id,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateJobResponse struct{}"
	}

	return strings.Join([]string{"CreateJobResponse", string(data)}, " ")
}
