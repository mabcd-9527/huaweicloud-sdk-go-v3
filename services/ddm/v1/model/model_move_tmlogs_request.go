package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MoveTmlogsRequest 请求体
type MoveTmlogsRequest struct {

	// 节点ID列表
	NodeIds []string `json:"node_ids"`
}

func (o MoveTmlogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MoveTmlogsRequest struct{}"
	}

	return strings.Join([]string{"MoveTmlogsRequest", string(data)}, " ")
}
