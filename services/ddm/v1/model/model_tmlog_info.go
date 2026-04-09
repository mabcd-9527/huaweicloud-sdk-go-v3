package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TmlogInfo struct {

	// 节点ID
	NodeId *string `json:"node_id,omitempty"`

	// 压缩文件列表
	FileList *[]string `json:"file_list,omitempty"`

	// TMLOG文件大小单位（MB）。
	TmlogSize float32 `json:"tmlog_size,omitempty"`
}

func (o TmlogInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TmlogInfo struct{}"
	}

	return strings.Join([]string{"TmlogInfo", string(data)}, " ")
}
