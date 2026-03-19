package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateCcRuleResponseBodyAction 请求次数限制到达后采取的防护动作
type UpdateCcRuleResponseBodyAction struct {

	// 动作类型：  - captcha：人机验证，阻断后用户需要输入正确的验证码，恢复正确的访问页面。  - block：阻断。   - log: 仅记录   - dynamic_block: 上一个限速周期内，请求频率超过“限速频率”将被阻断，那么在下一个限速周期内，请求频率超过“放行频率”将被阻断。注：只有当cc防护规则模式为高级模式时才支持设置dynamic_block防护动作。
	Category UpdateCcRuleResponseBodyActionCategory `json:"category"`

	Detail *CcrulesListInfoActionDetail `json:"detail,omitempty"`
}

func (o UpdateCcRuleResponseBodyAction) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCcRuleResponseBodyAction struct{}"
	}

	return strings.Join([]string{"UpdateCcRuleResponseBodyAction", string(data)}, " ")
}

type UpdateCcRuleResponseBodyActionCategory struct {
	value string
}

type UpdateCcRuleResponseBodyActionCategoryEnum struct {
	CAPTCHA       UpdateCcRuleResponseBodyActionCategory
	BLOCK         UpdateCcRuleResponseBodyActionCategory
	LOG           UpdateCcRuleResponseBodyActionCategory
	DYNAMIC_BLOCK UpdateCcRuleResponseBodyActionCategory
}

func GetUpdateCcRuleResponseBodyActionCategoryEnum() UpdateCcRuleResponseBodyActionCategoryEnum {
	return UpdateCcRuleResponseBodyActionCategoryEnum{
		CAPTCHA: UpdateCcRuleResponseBodyActionCategory{
			value: "captcha",
		},
		BLOCK: UpdateCcRuleResponseBodyActionCategory{
			value: "block",
		},
		LOG: UpdateCcRuleResponseBodyActionCategory{
			value: "log",
		},
		DYNAMIC_BLOCK: UpdateCcRuleResponseBodyActionCategory{
			value: "dynamic_block",
		},
	}
}

func (c UpdateCcRuleResponseBodyActionCategory) Value() string {
	return c.value
}

func (c UpdateCcRuleResponseBodyActionCategory) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateCcRuleResponseBodyActionCategory) UnmarshalJSON(b []byte) error {
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
