package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImageLocalResponse Response Object
type ListImageLocalResponse struct {

	// **参数解释** 符合所有筛选条件的本地镜像总记录数，用于分页计算总页数 **取值范围** 取值0-2147483647
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释** 包含查询到的本地镜像详情，每个元素对应一个本地镜像的完整信息 **取值范围** 数组长度0-limit（每页显示个数），元素结构符合ImageLocalInfo定义，数组为空表示无匹配结果
	DataList       *[]ImageLocalInfo `json:"data_list,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListImageLocalResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImageLocalResponse struct{}"
	}

	return strings.Join([]string{"ListImageLocalResponse", string(data)}, " ")
}
