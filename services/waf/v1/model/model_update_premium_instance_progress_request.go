package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePremiumInstanceProgressRequest Request Object
type UpdatePremiumInstanceProgressRequest struct {

	// **参数解释：** 独享模式域名Id，通过 查询独享模式域名列表(ListPremiumHost) 接口获取 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	HostId string `json:"host_id"`

	// **参数解释：** 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目ID。若需要查询当前用户所有企业项目绑定的资源信息，请传参all_granted_eps。 **约束限制：** 不涉及 **取值范围：**  - 0：代表default企业项目  - all_granted_eps：代表所有企业项目  - 其它企业项目ID：长度为36个字符 **默认取值：** 0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *AccessProgress `json:"body,omitempty"`
}

func (o UpdatePremiumInstanceProgressRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePremiumInstanceProgressRequest struct{}"
	}

	return strings.Join([]string{"UpdatePremiumInstanceProgressRequest", string(data)}, " ")
}
