package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEndpointResponse Response Object
type UpdateEndpointResponse struct {

	// 可见性：  - PRIVATE：私有  - PUBLIC：公共  默认为PRIVATE
	Visibility *string `json:"visibility,omitempty"`

	// Endpoint Id，32~36位的英文、数字、短横组合。
	Id string `json:"id"`

	// 一个Endpoint名称，只能包含中文、字母、数字、下划线、中划线、点、空格
	Name string `json:"name"`

	Type *EndpointType `json:"type"`

	Status *EndpointStatus `json:"status"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time"`

	// 更新时间
	UpdateTime *sdktime.SdkTime `json:"update_time"`

	Owner *User `json:"owner"`

	Cap *CapRef `json:"cap,omitempty"`

	// - **参数解释**：CustomImageId。 - **约束限制**：不涉及。 - **取值范围**：长度为[32,36]的英文字符、数字和中划线(-)的组合。 - **默认取值**：不涉及。
	CustomImageId *string `json:"custom_image_id,omitempty"`

	ReservedResource *ReservedResource `json:"reserved_resource,omitempty"`

	RayResource *RayResourceInfo `json:"ray_resource,omitempty"`

	// 失败状态时的错误编码
	ErrorCode *string `json:"error_code,omitempty"`

	// 失败状态时的错误信息
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 调用地址
	Urls *[]Url `json:"urls,omitempty"`

	// 引擎实例id列表
	BusinessEngineInstanceIds *[]string `json:"business_engine_instance_ids,omitempty"`

	TokensQuota *TokensQuota `json:"tokens_quota,omitempty"`

	RuntimeEnvType *RuntimeEnvType `json:"runtime_env_type,omitempty"`

	Config *EndpointConfigResponse `json:"config,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateEndpointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEndpointResponse struct{}"
	}

	return strings.Join([]string{"UpdateEndpointResponse", string(data)}, " ")
}
