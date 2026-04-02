package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type JobParameters struct {

	// 是否删除无tag的制品；默认为false，不删除无tag的制品。
	DeleteUntagged *bool `json:"delete_untagged,omitempty"`

	// 并行执行制品清理任务的工作者数量，最小值为1，最大值为5。
	Workers int32 `json:"workers"`

	// 是否模拟执行任务；默认值为false，非模拟运行。
	DryRun *bool `json:"dry_run,omitempty"`

	// 清理镜像层的个数。
	PurgedBlobs *int32 `json:"purged_blobs,omitempty"`

	// 清理镜像manifest的个数。
	PurgedManifests *int32 `json:"purged_manifests,omitempty"`

	// 释放的镜像存储空间大小(单位为字节)。
	FreedSpace *int32 `json:"freed_space,omitempty"`
}

func (o JobParameters) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobParameters struct{}"
	}

	return strings.Join([]string{"JobParameters", string(data)}, " ")
}
