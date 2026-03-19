package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateSqlLimitTaskNodeOption struct {

	// **参数解释**: 节点ID。 **约束限制**: 必须是当前实例的某一个节点ID。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	NodeId string `json:"node_id"`

	// **参数解释**: 该节点执行的SQL语句ID。 **约束限制**: 每个节点最多可以下发的SQL ID数量不可以超过10。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	SqlIds []string `json:"sql_ids"`
}

func (o CreateSqlLimitTaskNodeOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSqlLimitTaskNodeOption struct{}"
	}

	return strings.Join([]string{"CreateSqlLimitTaskNodeOption", string(data)}, " ")
}
