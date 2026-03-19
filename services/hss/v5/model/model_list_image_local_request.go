package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListImageLocalRequest Request Object
type ListImageLocalRequest struct {

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**: 偏移量：指定返回记录的开始位置 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值2000000 **默认取值**: 默认为0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 每页显示个数 **约束限制**: 不涉及 **取值范围**: 取值10-200 **默认取值**: 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释** 本地镜像的名称，用于模糊筛选指定名称的本地镜像列表 **约束限制** 支持部分匹配（如传入'web'可匹配所有名称含'web'的镜像），区分大小写 **取值范围** 字符长度1-256位，支持字母、数字、短横线、下划线、点号，禁止含@#$%等特殊字符 **默认取值** 无
	ImageName *string `json:"image_name,omitempty"`

	// **参数解释** 本地镜像的版本标识，用于筛选指定版本的本地镜像，需与image_name配合使用 **约束限制** 仅当指定image_name时传参有效，否则筛选条件不生效 **取值范围** 字符长度1-128位，支持字母、数字、短横线、下划线、点号、冒号 **默认取值** 无
	ImageVersion *string `json:"image_version,omitempty"`

	// **参数解释** 本地镜像的安全扫描状态，用于筛选指定扫描状态的镜像列表 **约束限制** 取值必须在指定范围内，否则返回空结果，区分大小写 **取值范围**   - unscan : 未扫描   - success : 扫描完成   - scanning : 扫描中   - failed : 扫描失败   - waiting_for_scan : 等待扫描 **默认取值** 无
	ScanStatus *string `json:"scan_status,omitempty"`

	// **参数解释** 本地镜像的存储来源类型，用于筛选不同来源的本地镜像 **约束限制** 取值必须在指定范围内，否则返回空结果，区分大小写 **取值范围**  - other_image : 非SWR镜像  - swr_image : SWR镜像 **默认取值** 无
	LocalImageType *string `json:"local_image_type,omitempty"`

	// **参数解释** 本地镜像的大小（单位字节），用于筛选指定大小的镜像（精确匹配） **约束限制** 仅支持精确匹配，如需范围筛选需结合业务层处理 **取值范围** 取值0-9223372036854775807（约9EB） **默认取值** 无
	ImageSize *int64 `json:"image_size,omitempty"`

	// **参数解释** 本地镜像版本最后更新时间的查询起始值（Unix时间戳，单位ms），与end_latest_update_time配合筛选时间范围 **时间格式** Unix时间戳（如1697509433000表示2023-10-16 10:23:53） **约束限制** 需与end_latest_update_time同时使用，且小于end_latest_update_time，否则筛选无效 **取值范围** 取值0-9223372036854775807 **默认取值** 无
	StartLatestUpdateTime *int64 `json:"start_latest_update_time,omitempty"`

	// **参数解释** 本地镜像版本最后更新时间的查询结束值（Unix时间戳，单位ms），与start_latest_update_time配合筛选时间范围 **时间格式** Unix时间戳（如1709973506292表示2024-03-08 15:18:26） **约束限制** 需与start_latest_update_time同时使用，且大于start_latest_update_time，否则筛选无效 **取值范围** 取值0-9223372036854775807 **默认取值** 无
	EndLatestUpdateTime *int64 `json:"end_latest_update_time,omitempty"`

	// **参数解释** 本地镜像最近一次扫描完成时间的查询起始值（Unix时间戳，单位ms），与end_latest_scan_time配合筛选时间范围 **时间格式** Unix时间戳（精确到毫秒） **约束限制** 仅对scan_status为success的镜像有效，需与end_latest_scan_time同时使用 **取值范围** 取值0-9223372036854775807 **默认取值** 无
	StartLatestScanTime *int64 `json:"start_latest_scan_time,omitempty"`

	// **参数解释** 本地镜像最近一次扫描完成时间的查询结束值（Unix时间戳，单位ms），与start_latest_scan_time配合筛选时间范围 **时间格式** Unix时间戳（精确到毫秒） **约束限制** 仅对scan_status为success的镜像有效，且需大于start_latest_scan_time **取值范围** 取值0-9223372036854775807 **默认取值** 无
	EndLatestScanTime *int64 `json:"end_latest_scan_time,omitempty"`

	// **参数解释** 用于筛选是否存在软件漏洞的本地镜像，true表示筛选有漏洞的镜像，false表示筛选无漏洞的镜像 **约束限制** 仅对scan_status为success的镜像有效，未扫描镜像不会被筛选 **取值范围** true（存在漏洞）、false（不存在漏洞） **默认取值** 无（不筛选漏洞状态）
	HasVul *bool `json:"has_vul,omitempty"`

	// **参数解释** 本地镜像所关联的云服务器名称，用于筛选关联指定服务器的本地镜像 **约束限制** 支持模糊匹配，区分大小写，仅对关联了服务器的镜像有效 **取值范围** 字符长度1-64位，支持中文、英文、数字、短横线、下划线，禁止含特殊字符 **默认取值** 无
	HostName *string `json:"host_name,omitempty"`

	// **参数解释** 本地镜像所关联的云服务器唯一标识（ECS实例ID），用于精准筛选关联指定服务器的本地镜像 **约束限制** 精确匹配，仅对关联了该服务器的镜像有效 **取值范围** 字符长度1-64位，支持字母、数字、短横线 **默认取值** 无
	HostId *string `json:"host_id,omitempty"`

	// **参数解释** 本地镜像所关联服务器的公网或私网IP地址，用于筛选关联指定IP服务器的本地镜像 **约束限制** 支持IPv4格式，精确匹配，多个IP需通过业务层分批查询 **取值范围** 符合IPv4格式的字符串（如 **默认取值** 无
	HostIp *string `json:"host_ip,omitempty"`

	// **参数解释** 本地镜像所关联的容器唯一标识（Docker容器ID），用于精准筛选关联指定容器的本地镜像 **约束限制** 精确匹配，仅对关联了容器的镜像有效 **取值范围** 字符长度1-64位，支持字母、数字、短横线、下划线 **默认取值** 无
	ContainerId *string `json:"container_id,omitempty"`

	// **参数解释** 本地镜像所关联的容器名称，用于筛选关联指定名称容器的本地镜像 **约束限制** 支持模糊匹配，区分大小写，仅对关联了容器的镜像有效 **取值范围** 字符长度1-64位，支持字母、数字、短横线、下划线、点号 **默认取值** 无
	ContainerName *string `json:"container_name,omitempty"`

	// **参数解释** 本地镜像所关联的Kubernetes Pod唯一标识，用于精准筛选关联指定Pod的本地镜像 **约束限制** 精确匹配，仅对K8s环境中关联了Pod的镜像有效 **取值范围** 字符长度1-64位，支持字母、数字、短横线 **默认取值** 无
	PodId *string `json:"pod_id,omitempty"`

	// **参数解释** 本地镜像所关联的Kubernetes Pod名称，用于筛选关联指定名称Pod的本地镜像 **约束限制** 支持模糊匹配，区分大小写，仅对K8s环境中关联了Pod的镜像有效 **取值范围** 字符长度1-63位，支持字母、数字、短横线，不能以短横线开头或结尾 **默认取值** 无
	PodName *string `json:"pod_name,omitempty"`

	// **参数解释** 本地镜像中部署的应用软件名称（如Nginx、MySQL），用于筛选包含指定应用的本地镜像 **约束限制** 支持模糊匹配，区分大小写，仅对已识别应用的镜像有效 **取值范围** 字符长度1-64位，支持中文、英文、数字、短横线、下划线 **默认取值** 无
	AppName *string `json:"app_name,omitempty"`

	// **参数解释** 用于筛选是否关联了容器的本地镜像 **取值范围**: - true：关联容器的镜像 - false：未关联容器的镜像 **默认取值** 无（不筛选容器关联状态）
	HasContainer *bool `json:"has_container,omitempty"`
}

func (o ListImageLocalRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListImageLocalRequest struct{}"
	}

	return strings.Join([]string{"ListImageLocalRequest", string(data)}, " ")
}
