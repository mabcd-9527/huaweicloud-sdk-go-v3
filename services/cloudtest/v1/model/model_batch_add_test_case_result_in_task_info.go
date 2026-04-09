package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchAddTestCaseResultInTaskInfo struct {
	Result *AddTestCaseResultInfo `json:"result,omitempty"`

	// 任务uri
	TaskUri *string `json:"task_uri,omitempty"`

	// 测试套结果URI
	TaskResultUri *string `json:"task_result_uri,omitempty"`

	// 用例uri列表
	TestCaseUris *[]string `json:"test_case_uris,omitempty"`

	// 是否异步执行
	IsAsyn *bool `json:"isAsyn,omitempty"`
}

func (o BatchAddTestCaseResultInTaskInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddTestCaseResultInTaskInfo struct{}"
	}

	return strings.Join([]string{"BatchAddTestCaseResultInTaskInfo", string(data)}, " ")
}
