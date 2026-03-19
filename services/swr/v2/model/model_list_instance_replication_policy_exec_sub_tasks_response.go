package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceReplicationPolicyExecSubTasksResponse Response Object
type ListInstanceReplicationPolicyExecSubTasksResponse struct {

	// 镜像同步策略执行记录子任务列表
	SubTasks *[]SubtaskDetail `json:"sub_tasks,omitempty"`

	// 镜像同步策略执行记录子任务总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListInstanceReplicationPolicyExecSubTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceReplicationPolicyExecSubTasksResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceReplicationPolicyExecSubTasksResponse", string(data)}, " ")
}
