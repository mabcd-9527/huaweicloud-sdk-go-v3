package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIpGroupRequest Request Object
type ShowIpGroupRequest struct {

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id。若需要查询当前用户所有企业项目绑定的资源信息，请传参all_granted_eps。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释：** ip地址组id，可从查询地址组列表(ListIpGroup)接口中获取 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Id string `json:"id"`
}

func (o ShowIpGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIpGroupRequest struct{}"
	}

	return strings.Join([]string{"ShowIpGroupRequest", string(data)}, " ")
}
