package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCustomRulesRequest Request Object
type ListCustomRulesRequest struct {

	// **参数解释：** 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目ID。若需要查询当前用户所有企业项目绑定的资源信息，请传参all_granted_eps。 **约束限制：** 不涉及 **取值范围：**  - 0：代表default企业项目  - all_granted_eps：代表所有企业项目  - 其它企业项目ID：长度为36个字符 **默认取值：** 0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释：** 防护策略id，您可以通过调用查询防护策略列表（ListPolicy）获取策略id **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	PolicyId string `json:"policy_id"`

	// **参数解释：** 偏移量，表示查询该偏移量之后的记录，与参数limit一起使用 **约束限制：** 不涉及 **取值范围：** [0, 65535] **默认取值：** 不涉及
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 查询返回记录的数量限制，与参数offset一起使用，如果offset未设置值，则limit无效 **约束限制：** 不涉及 **取值范围：** [1, 65535] **默认取值：** 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：** 当前页码，与参数pagesize一起使用 **约束限制：** 不涉及 **取值范围：** [1, 记录数/pagesize] **默认取值：** 1
	Page *int32 `json:"page,omitempty"`

	// **参数解释：** 每页大小，与参数page一起使用 **约束限制：** 不涉及 **取值范围：** [0, 2147483647] **默认取值：** 1000
	Pagesize *int32 `json:"pagesize,omitempty"`
}

func (o ListCustomRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCustomRulesRequest struct{}"
	}

	return strings.Join([]string{"ListCustomRulesRequest", string(data)}, " ")
}
