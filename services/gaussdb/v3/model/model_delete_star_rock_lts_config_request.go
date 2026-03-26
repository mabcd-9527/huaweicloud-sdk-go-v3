package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteStarRockLtsConfigRequest Request Object
type DeleteStarRockLtsConfigRequest struct {

	// **参数解释**：              请求语言类型。  **约束限制**：  不涉及。  **取值范围**： - en-us - zh-cn  **默认值**：  en-us。
	XLanguage DeleteStarRockLtsConfigRequestXLanguage `json:"X-Language"`

	Body *HtapDeleteLtsConfigRequestBody `json:"body,omitempty"`
}

func (o DeleteStarRockLtsConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteStarRockLtsConfigRequest struct{}"
	}

	return strings.Join([]string{"DeleteStarRockLtsConfigRequest", string(data)}, " ")
}

type DeleteStarRockLtsConfigRequestXLanguage struct {
	value string
}

type DeleteStarRockLtsConfigRequestXLanguageEnum struct {
	ZH_CN DeleteStarRockLtsConfigRequestXLanguage
	EN_US DeleteStarRockLtsConfigRequestXLanguage
}

func GetDeleteStarRockLtsConfigRequestXLanguageEnum() DeleteStarRockLtsConfigRequestXLanguageEnum {
	return DeleteStarRockLtsConfigRequestXLanguageEnum{
		ZH_CN: DeleteStarRockLtsConfigRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: DeleteStarRockLtsConfigRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c DeleteStarRockLtsConfigRequestXLanguage) Value() string {
	return c.value
}

func (c DeleteStarRockLtsConfigRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteStarRockLtsConfigRequestXLanguage) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
