package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateUserRequestV3 创建账号的请求对象
type CreateUserRequestV3 struct {

	// **参数解释**：  实例账号名称。  **约束限制**：  - 长度为1-32个字符。 - 必须以字母开头。 - 可以包含字母、数字、下划线，不能包含其它特殊字符。  **取值范围**：  不涉及。  **默认取值**：    不涉及。
	Name string `json:"name"`

	// **参数解释**：  实例账号密码。  **约束限制**：  - 长度为8~32个字符。 - 至少包含三种字符组合：小写字母、大写字母、数字、特殊字符 ~ ! @ # % ^ * - _ + ? - 不能使用简单、强度不够、容易被猜测的密码。 - 不能与账号名称或者倒序的账号名称相同。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Password string `json:"password"`

	// **参数解释**：  DDM实例账号的基础权限。  **约束限制**：  无  **取值范围**：  CREATE、DROP、ALTER、INDEX、INSERT、DELETE、UPDATE、SELECT的任意集合  **默认取值**：  不涉及。
	BaseAuthority []CreateUserRequestV3BaseAuthority `json:"base_authority"`

	// **参数解释**：  实例账号的描述信息。  **约束限制**：  - 长度不超过256个字符。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**：  DDM实例账号的密码有效期。  **约束限制**：  不涉及。   **取值范围**：  取值范围为0-65535的整数，单位为天。  0与空表示密码永不过期。  **默认取值**：  默认值为空，永不过期。
	PasswordLifetime *int64 `json:"password_lifetime,omitempty"`

	// 关联的逻辑库的集合。 databases字段可以省略，即创建用户时可以不关联逻辑库。
	Databases *[]CreateUserRelatedLogicDbV3 `json:"databases,omitempty"`
}

func (o CreateUserRequestV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUserRequestV3 struct{}"
	}

	return strings.Join([]string{"CreateUserRequestV3", string(data)}, " ")
}

type CreateUserRequestV3BaseAuthority struct {
	value string
}

type CreateUserRequestV3BaseAuthorityEnum struct {
	CREATE CreateUserRequestV3BaseAuthority
	DROP   CreateUserRequestV3BaseAuthority
	ALTER  CreateUserRequestV3BaseAuthority
	INDEX  CreateUserRequestV3BaseAuthority
	INSERT CreateUserRequestV3BaseAuthority
	DELETE CreateUserRequestV3BaseAuthority
	UPDATE CreateUserRequestV3BaseAuthority
	SELECT CreateUserRequestV3BaseAuthority
}

func GetCreateUserRequestV3BaseAuthorityEnum() CreateUserRequestV3BaseAuthorityEnum {
	return CreateUserRequestV3BaseAuthorityEnum{
		CREATE: CreateUserRequestV3BaseAuthority{
			value: "CREATE",
		},
		DROP: CreateUserRequestV3BaseAuthority{
			value: "DROP",
		},
		ALTER: CreateUserRequestV3BaseAuthority{
			value: "ALTER",
		},
		INDEX: CreateUserRequestV3BaseAuthority{
			value: "INDEX",
		},
		INSERT: CreateUserRequestV3BaseAuthority{
			value: "INSERT",
		},
		DELETE: CreateUserRequestV3BaseAuthority{
			value: "DELETE",
		},
		UPDATE: CreateUserRequestV3BaseAuthority{
			value: "UPDATE",
		},
		SELECT: CreateUserRequestV3BaseAuthority{
			value: "SELECT",
		},
	}
}

func (c CreateUserRequestV3BaseAuthority) Value() string {
	return c.value
}

func (c CreateUserRequestV3BaseAuthority) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateUserRequestV3BaseAuthority) UnmarshalJSON(b []byte) error {
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
