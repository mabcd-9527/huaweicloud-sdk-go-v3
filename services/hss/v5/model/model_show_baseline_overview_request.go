package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBaselineOverviewRequest Request Object
type ShowBaselineOverviewRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 策略组ID。当传此字段时，返回当前策略组ID下的数据。当不传此字段时，返回当前企业项目下所有数据。 **约束限制**: 不涉及 **取值范围**: 字符长度0-256位 **默认取值**: 不涉及
	GroupId *string `json:"group_id,omitempty"`

	// **参数解释** 统计类型，不同的查询类型返回不同的统计数据，为空表示返回所有数据 **约束限制** 不涉及 **取值范围** - weakpwd: 查询主机弱口令检测统计数据 - baseline_passrate: 查询基线检查通过率、主机配置基线通过率、基线风险统计、未通过检查项/已检查项 - host_risk_top_5: 查询主机配置风险TO5 - host_num: 未通过主机/主机总数  **默认取值** 不涉及
	StatisticsType *string `json:"statistics_type,omitempty"`
}

func (o ShowBaselineOverviewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBaselineOverviewRequest struct{}"
	}

	return strings.Join([]string{"ShowBaselineOverviewRequest", string(data)}, " ")
}
