package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateJobResponse Response Object
type UpdateJobResponse struct {

	// 作业Id
	Id *string `json:"id,omitempty"`

	// 作业版本Id
	VersionId      *string `json:"version_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateJobResponse struct{}"
	}

	return strings.Join([]string{"UpdateJobResponse", string(data)}, " ")
}
