package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IdNameEntry struct {

	// 资源id
	Id *string `json:"id,omitempty"`

	// 资源名称
	Name *string `json:"name,omitempty"`

	// 引擎实例IP
	ServiceIp *string `json:"service_ip,omitempty"`
}

func (o IdNameEntry) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IdNameEntry struct{}"
	}

	return strings.Join([]string{"IdNameEntry", string(data)}, " ")
}
