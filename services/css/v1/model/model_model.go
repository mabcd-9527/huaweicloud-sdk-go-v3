package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Model **参数解释**： 模型 **取值范围**： 不涉及
type Model struct {

	// **参数解释**： 模型id **取值范围**： 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**： 模型名字 **取值范围**： 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**： 模型类型 **取值范围**： 不涉及
	DatastoreType *string `json:"datastore_type,omitempty"`

	// **参数解释**： 模型版本 **取值范围**： 不涉及
	DatastoreVersion *string `json:"datastore_version,omitempty"`

	// **参数解释**： 是否是文本模型 **取值范围**： 不涉及
	IsTextModel *string `json:"is_text_model,omitempty"`

	// **参数解释**： 模型版本id **取值范围**： 不涉及
	ModelVersionId *string `json:"model_version_id,omitempty"`

	// **参数解释**： 模型描述 **取值范围**： 不涉及
	Desc *string `json:"desc,omitempty"`

	// **参数解释**： 模型语言 **取值范围**： 不涉及
	Language *string `json:"language,omitempty"`

	// **参数解释**： 模型规格 **取值范围**： 不涉及
	ArchType *string `json:"arch_type,omitempty"`
}

func (o Model) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Model struct{}"
	}

	return strings.Join([]string{"Model", string(data)}, " ")
}
