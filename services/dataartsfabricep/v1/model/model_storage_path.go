package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// StoragePath 存储路径
type StoragePath struct {

	// 日志存储类型，可选值如下： OBS：日志存储在OBS CONTAINER：日志存储在容器 LTS：日志存储在LTS
	Type *StoragePathType `json:"type,omitempty"`

	// 路径
	Path *string `json:"path,omitempty"`
}

func (o StoragePath) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StoragePath struct{}"
	}

	return strings.Join([]string{"StoragePath", string(data)}, " ")
}

type StoragePathType struct {
	value string
}

type StoragePathTypeEnum struct {
	OBS       StoragePathType
	CONTAINER StoragePathType
	LTS       StoragePathType
}

func GetStoragePathTypeEnum() StoragePathTypeEnum {
	return StoragePathTypeEnum{
		OBS: StoragePathType{
			value: "OBS",
		},
		CONTAINER: StoragePathType{
			value: "CONTAINER",
		},
		LTS: StoragePathType{
			value: "LTS",
		},
	}
}

func (c StoragePathType) Value() string {
	return c.value
}

func (c StoragePathType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StoragePathType) UnmarshalJSON(b []byte) error {
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
