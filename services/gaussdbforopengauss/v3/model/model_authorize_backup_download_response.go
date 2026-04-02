package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AuthorizeBackupDownloadResponse Response Object
type AuthorizeBackupDownloadResponse struct {

	// **参数解释**: 文件所在桶名。 **取值范围**: 不涉及。
	Bucket *string `json:"bucket,omitempty"`

	// **参数解释**: 通过OBS Browser+下载备份文件的路径名称。
	FilePaths      *[]string `json:"file_paths,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o AuthorizeBackupDownloadResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuthorizeBackupDownloadResponse struct{}"
	}

	return strings.Join([]string{"AuthorizeBackupDownloadResponse", string(data)}, " ")
}
