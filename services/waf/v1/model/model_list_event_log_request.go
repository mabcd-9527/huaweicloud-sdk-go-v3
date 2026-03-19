package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEventLogRequest Request Object
type ListEventLogRequest struct {

	// **参数解释：** 分页查询时，返回第几页数据 **约束限制：** 不涉及 **取值范围：** page参数的实际有效范围取决于总数据量和pagesize的取值，不能大于总页数 **默认取值：** 1
	Page *int32 `json:"page,omitempty"`

	// **参数解释：** 分页查询时，每页包含的结果条数 **约束限制：** 不涉及 **取值范围：** [0, 总数据量] **默认取值：** 10
	Pagesize *int32 `json:"pagesize,omitempty"`
}

func (o ListEventLogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEventLogRequest struct{}"
	}

	return strings.Join([]string{"ListEventLogRequest", string(data)}, " ")
}
