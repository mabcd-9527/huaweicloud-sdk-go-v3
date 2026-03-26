package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EndpointInput 创建Endpoint的输入参数
type EndpointInput struct {

	// 一个Endpoint名称，只能包含中文、字母、数字、下划线、中划线、点、空格
	Name string `json:"name"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	Type *EndpointType `json:"type"`

	ReservedResource *ReservedResource `json:"reserved_resource,omitempty"`

	RayResource *RayResourceInput `json:"ray_resource,omitempty"`

	Cap *CapRef `json:"cap,omitempty"`

	Config *EndpointConfig `json:"config,omitempty"`

	CacheList *[]CacheConfig `json:"cache_list,omitempty"`

	// **参数解释**：是否开启公网访问。 **约束限制**：不涉及。 **取值范围**：开启true，关闭false。 **默认取值**：false。
	PublicAccess *bool `json:"public_access,omitempty"`

	// 自定义镜像ID
	CustomImageId *string `json:"custom_image_id,omitempty"`
}

func (o EndpointInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EndpointInput struct{}"
	}

	return strings.Join([]string{"EndpointInput", string(data)}, " ")
}
