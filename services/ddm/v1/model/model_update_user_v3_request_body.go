package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateUserV3RequestBody 修改账号的请求对象
type UpdateUserV3RequestBody struct {

	// **参数解释**：  DDM实例账号的基础权限。  **约束限制**：  无  **取值范围**：  CREATE、DROP、ALTER、INDEX、INSERT、DELETE、UPDATE、SELECT的任意集合  **默认取值**：  不涉及。
	BaseAuthority *[]UpdateUserV3RequestBodyBaseAuthority `json:"base_authority,omitempty"`

	// **参数解释**：  实例账号的描述信息。  **约束限制**：  - 长度不超过256个字符。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**：  DDM实例账号的密码有效期。  **约束限制**：  不涉及。   **取值范围**：  取值范围为0-65535的整数，单位为天。  0与空表示密码永不过期。  **默认取值**：  默认值为空，永不过期。
	PasswordLifetime *int64 `json:"password_lifetime,omitempty"`

	// 关联的逻辑库的集合。 databases字段可以省略，即创建用户时可以不关联逻辑库。
	Databases *[]UpdateUserRelatedLogicDbV3 `json:"databases,omitempty"`
}

func (o UpdateUserV3RequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUserV3RequestBody struct{}"
	}

	return strings.Join([]string{"UpdateUserV3RequestBody", string(data)}, " ")
}

type UpdateUserV3RequestBodyBaseAuthority struct {
	value string
}

type UpdateUserV3RequestBodyBaseAuthorityEnum struct {
	CREATE UpdateUserV3RequestBodyBaseAuthority
	DROP   UpdateUserV3RequestBodyBaseAuthority
	ALTER  UpdateUserV3RequestBodyBaseAuthority
	INDEX  UpdateUserV3RequestBodyBaseAuthority
	INSERT UpdateUserV3RequestBodyBaseAuthority
	DELETE UpdateUserV3RequestBodyBaseAuthority
	UPDATE UpdateUserV3RequestBodyBaseAuthority
	SELECT UpdateUserV3RequestBodyBaseAuthority
}

func GetUpdateUserV3RequestBodyBaseAuthorityEnum() UpdateUserV3RequestBodyBaseAuthorityEnum {
	return UpdateUserV3RequestBodyBaseAuthorityEnum{
		CREATE: UpdateUserV3RequestBodyBaseAuthority{
			value: "CREATE",
		},
		DROP: UpdateUserV3RequestBodyBaseAuthority{
			value: "DROP",
		},
		ALTER: UpdateUserV3RequestBodyBaseAuthority{
			value: "ALTER",
		},
		INDEX: UpdateUserV3RequestBodyBaseAuthority{
			value: "INDEX",
		},
		INSERT: UpdateUserV3RequestBodyBaseAuthority{
			value: "INSERT",
		},
		DELETE: UpdateUserV3RequestBodyBaseAuthority{
			value: "DELETE",
		},
		UPDATE: UpdateUserV3RequestBodyBaseAuthority{
			value: "UPDATE",
		},
		SELECT: UpdateUserV3RequestBodyBaseAuthority{
			value: "SELECT",
		},
	}
}

func (c UpdateUserV3RequestBodyBaseAuthority) Value() string {
	return c.value
}

func (c UpdateUserV3RequestBodyBaseAuthority) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateUserV3RequestBodyBaseAuthority) UnmarshalJSON(b []byte) error {
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
