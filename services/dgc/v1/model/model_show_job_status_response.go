package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowJobStatusResponse Response Object
type ShowJobStatusResponse struct {

	// 作业名称。
	Name *string `json:"name,omitempty"`

	// 作业状态： - STARTING：启动中 - NORMAL：正常 - EXCEPTION：异常 - STOPPING： 停止中 - STOPPED：停止 - PAUSE: 暂停 - ABNORMAL: 运行异常/失败 - DISABLE: 节点禁用 - OVERLOAD: 繁忙 - INIT: 初始化
	Status *ShowJobStatusResponseStatus `json:"status,omitempty"`

	// 开始时间。
	Starttime *string `json:"starttime,omitempty"`

	// 结束时间。
	EndTime *string `json:"endTime,omitempty"`

	// 状态最后更新时间。
	LastUpdateTime *string `json:"lastUpdateTime,omitempty"`

	Nodes *[]RealTimeNodeStatus `json:"nodes,omitempty"`

	// 创建人。
	CreateUser *string `json:"createUser,omitempty"`

	// 最后修改人。
	LastModifiyUser *string `json:"lastModifiyUser,omitempty"`
	HttpStatusCode  int     `json:"-"`
}

func (o ShowJobStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowJobStatusResponse", string(data)}, " ")
}

type ShowJobStatusResponseStatus struct {
	value string
}

type ShowJobStatusResponseStatusEnum struct {
	STARTING  ShowJobStatusResponseStatus
	NORMAL    ShowJobStatusResponseStatus
	EXCEPTION ShowJobStatusResponseStatus
	STOPPING  ShowJobStatusResponseStatus
	STOPPED   ShowJobStatusResponseStatus
	PAUSE     ShowJobStatusResponseStatus
	ABNORMAL  ShowJobStatusResponseStatus
	DISABLE   ShowJobStatusResponseStatus
	OVERLOAD  ShowJobStatusResponseStatus
	INIT      ShowJobStatusResponseStatus
}

func GetShowJobStatusResponseStatusEnum() ShowJobStatusResponseStatusEnum {
	return ShowJobStatusResponseStatusEnum{
		STARTING: ShowJobStatusResponseStatus{
			value: "STARTING",
		},
		NORMAL: ShowJobStatusResponseStatus{
			value: "NORMAL",
		},
		EXCEPTION: ShowJobStatusResponseStatus{
			value: "EXCEPTION",
		},
		STOPPING: ShowJobStatusResponseStatus{
			value: "STOPPING",
		},
		STOPPED: ShowJobStatusResponseStatus{
			value: "STOPPED",
		},
		PAUSE: ShowJobStatusResponseStatus{
			value: "PAUSE",
		},
		ABNORMAL: ShowJobStatusResponseStatus{
			value: "ABNORMAL",
		},
		DISABLE: ShowJobStatusResponseStatus{
			value: "DISABLE",
		},
		OVERLOAD: ShowJobStatusResponseStatus{
			value: "OVERLOAD",
		},
		INIT: ShowJobStatusResponseStatus{
			value: "INIT",
		},
	}
}

func (c ShowJobStatusResponseStatus) Value() string {
	return c.value
}

func (c ShowJobStatusResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowJobStatusResponseStatus) UnmarshalJSON(b []byte) error {
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
