package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BacthCreateWhiteBlackIpRuleRequestBody 创建黑白名单规则body
type BacthCreateWhiteBlackIpRuleRequestBody struct {

	// **参数解释：** 规则名称 **约束限制：** 长度范围：[1, 256] **取值范围：** 不涉及 **默认取值：** 不涉及
	Name string `json:"name"`

	// **参数解释：** ip地址，需要输入标准的ip地址或地址段，例如：42.123.120.66或42.123.120.0/16 **约束限制：** 参数“addr”和“ip_group_id”必须存在一个，同时存在时以参数“addr”为准 **取值范围：** 不涉及 **默认取值：** 不涉及
	Addr *string `json:"addr,omitempty"`

	// 黑白名单规则描述
	Description *string `json:"description,omitempty"`

	// **参数解释：** 防护动作 **约束限制：** 不涉及 **取值范围：**  - 0 拦截  - 1 放行   - 2 仅记录 **默认取值：** 不涉及
	White int32 `json:"white"`

	// **参数解释：** Ip地址组id，可在控制台中对象管理->地址组管理中添加 **约束限制：** 参数“addr”和“ip_group_id”必须存在一个，同时存在时以参数“addr”为准 **取值范围：** 不涉及 **默认取值：** 不涉及
	IpGroupId *string `json:"ip_group_id,omitempty"`

	// **参数解释：** 生效模式 **约束限制：** 不涉及 **取值范围：** - permanent 立即生效 - customize 自定义生效 **默认取值：** permanent
	TimeMode *BacthCreateWhiteBlackIpRuleRequestBodyTimeMode `json:"time_mode,omitempty"`

	// 规则生效开始时间，生效模式为自定义时，此字段才有效，请输入时间戳
	Start *int32 `json:"start,omitempty"`

	// 规则生效结束时间，生效模式为自定义时，此字段才有效，请输入时间戳
	Terminal *int32 `json:"terminal,omitempty"`

	// **参数解释：** 添加规则的策略id列表。策略id从\"查询防护策略列表\"(ListPolicy)接口获取，多个策略之间用“,”隔开 **约束限制：** 不能为空 **取值范围：** 不涉及 **默认取值：** 不涉及
	PolicyIds []string `json:"policy_ids"`
}

func (o BacthCreateWhiteBlackIpRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BacthCreateWhiteBlackIpRuleRequestBody struct{}"
	}

	return strings.Join([]string{"BacthCreateWhiteBlackIpRuleRequestBody", string(data)}, " ")
}

type BacthCreateWhiteBlackIpRuleRequestBodyTimeMode struct {
	value string
}

type BacthCreateWhiteBlackIpRuleRequestBodyTimeModeEnum struct {
	PERMANENT BacthCreateWhiteBlackIpRuleRequestBodyTimeMode
	CUSTOMIZE BacthCreateWhiteBlackIpRuleRequestBodyTimeMode
}

func GetBacthCreateWhiteBlackIpRuleRequestBodyTimeModeEnum() BacthCreateWhiteBlackIpRuleRequestBodyTimeModeEnum {
	return BacthCreateWhiteBlackIpRuleRequestBodyTimeModeEnum{
		PERMANENT: BacthCreateWhiteBlackIpRuleRequestBodyTimeMode{
			value: "permanent",
		},
		CUSTOMIZE: BacthCreateWhiteBlackIpRuleRequestBodyTimeMode{
			value: "customize",
		},
	}
}

func (c BacthCreateWhiteBlackIpRuleRequestBodyTimeMode) Value() string {
	return c.value
}

func (c BacthCreateWhiteBlackIpRuleRequestBodyTimeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BacthCreateWhiteBlackIpRuleRequestBodyTimeMode) UnmarshalJSON(b []byte) error {
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
