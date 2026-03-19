package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIpGroupRequest Request Object
type ListIpGroupRequest struct {

	// **参数解释：** 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目ID。若需要查询当前用户所有企业项目绑定的资源信息，请传参all_granted_eps。 **约束限制：** 不涉及 **取值范围：**  - 0：代表default企业项目  - all_granted_eps：代表所有企业项目  - 其它企业项目ID：长度为36个字符 **默认取值：** 0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 页码，默认值为1
	Page *int32 `json:"page,omitempty"`

	// 每页的条数，单页条数限制100，默认值为10
	Pagesize *int32 `json:"pagesize,omitempty"`

	// ip地址组名称，支持模糊查询
	Name *string `json:"name,omitempty"`

	// 单个ip地址或ip段，传入该参数将查询包含传入的ip地址或ip段的地址组
	Ip *string `json:"ip,omitempty"`
}

func (o ListIpGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIpGroupRequest struct{}"
	}

	return strings.Join([]string{"ListIpGroupRequest", string(data)}, " ")
}
