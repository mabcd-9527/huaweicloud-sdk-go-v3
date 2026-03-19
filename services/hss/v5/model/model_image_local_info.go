package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImageLocalInfo struct {

	// **参数解释** 镜像名称 **取值范围** 字符长度0-256位
	ImageName *string `json:"image_name,omitempty"`

	// **参数解释** 本地镜像的唯一标识，用于后续查询镜像详情、执行扫描等操作 **取值范围** 字符长度1-64位，支持字母、数字、短横线，符合UUID格式
	ImageId *string `json:"image_id,omitempty"`

	// **参数解释** 本地镜像的加密摘要（SHA-256算法），用于唯一标识镜像内容，避免篡改 **取值范围** 字符长度64-128位，以'sha256:'开头，后跟十六进制字符串（如sha256:ce0b5d91b072730d0bc9518f11efd07eb7fdb9f43251e11a96cab5b1918b7044）
	ImageDigest *string `json:"image_digest,omitempty"`

	// **参数解释** 镜像版本 **取值范围** 字符长度0-256位
	ImageVersion *string `json:"image_version,omitempty"`

	// **参数解释** 本地镜像的存储来源类型，标识镜像是否来自华为云SWR仓库 **取值范围** swr_image：华为云SWR仓库镜像 other_image：非SWR仓库镜像
	LocalImageType *string `json:"local_image_type,omitempty"`

	// **参数解释** 本地镜像的安全扫描状态，反映当前镜像是否完成安全检测 **取值范围** 扫描状态，包含如下：   - unscan：未扫描   - success：扫描完成   - scanning：正在扫描   - failed：扫描失败   - waiting：等待扫描
	ScanStatus *string `json:"scan_status,omitempty"`

	// **参数解释** 本地镜像的实际存储大小，单位为字节（B） **取值范围** 取值0-9223372036854775807（约9EB）
	ImageSize *int64 `json:"image_size,omitempty"`

	// **参数解释** 本地镜像版本的最后更新时间，即镜像创建或更新的时间戳 **取值范围** Unix时间戳（单位ms），取值0-9223372036854775807
	LatestUpdateTime *int64 `json:"latest_update_time,omitempty"`

	// **参数解释** 本地镜像最近一次完成安全扫描的时间戳，未扫描时该字段可能为0或空 **取值范围** Unix时间戳（单位ms），取值0-9223372036854775807，未扫描时为0
	LatestScanTime *int64 `json:"latest_scan_time,omitempty"`

	// **参数解释** 本地镜像中检测到的软件漏洞总数，包含高、中、低危漏洞 **取值范围** 取值0-9223372036854775807
	VulNum *int64 `json:"vul_num,omitempty"`

	// **参数解释** 本地镜像在安全基线扫描中未通过的检查项数量，反映镜像配置合规性 **取值范围** 取值0-9223372036854775807
	UnsafeSettingNum *int64 `json:"unsafe_setting_num,omitempty"`

	// **参数解释** 本地镜像中检测到的恶意文件（如病毒、木马）总数 **取值范围** 取值0-9223372036854775807
	MaliciousFileNum *int64 `json:"malicious_file_num,omitempty"`

	// **参数解释** 当前本地镜像所关联的云服务器总数 **取值范围** 取值0-9223372036854775807
	HostNum *int64 `json:"host_num,omitempty"`

	// **参数解释** 当前本地镜像所创建或关联的容器总数 **取值范围** 取值0-9223372036854775807
	ContainerNum *int64 `json:"container_num,omitempty"`

	// **参数解释** 本地镜像中包含的软件组件（如依赖库、应用程序）总数 **取值范围** 取值0-9223372036854775807
	ComponentNum *int64 `json:"component_num,omitempty"`

	// **参数解释** 当scan_status为failed时，该字段说明扫描失败的具体原因，未失败时为空字符串 **取值范围** - unknown_error：未知错误 - failed_to_match_agent：对应主机未开启容器版防护或agent离线 - create_container_failed：创建容器失败      - get_container_info_failed：获取容器信息失败 - docker_offline：docker引擎不在线 - get_docker_root_failed：获取容器根文件系统失败 - image_not_exist_or_docker_api_fault：镜像不存在或docker接口错误 - huge_image：超大镜像 - docker_root_in_nfs：容器根目录位于网络挂载 - response_timed_out：响应超时
	ScanFailedDesc *string `json:"scan_failed_desc,omitempty"`

	// **参数解释** 根据镜像的漏洞、基线违规、恶意文件情况综合评定的风险等级 **取值范围** - Security：安全 - Low：低危 - Medium：中危 - High：高危
	SeverityLevel *string `json:"severity_level,omitempty"`

	// **参数解释** 应用防护事件所属云服务器的名称，用于标识事件来源主机 **取值范围** 字符长度1-64位，支持中文、英文、数字、短横线、下划线，符合华为云ECS命名规范
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**: 主机id **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	HostId *string `json:"host_id,omitempty"`

	// **参数解释** 本地镜像所在服务器上安装的HSS Agent唯一标识，用于关联Agent相关操作 **取值范围** 字符长度1-128位，支持字母、数字、短横线、下划线
	AgentId *string `json:"agent_id,omitempty"`

	// **参数解释** 说明本地镜像无法进行安全扫描的具体原因（如镜像格式不支持、权限不足等），为空表示支持扫描 **取值范围** 字符长度0-1024位，支持中文、英文、数字、常用标点符号
	NonScanReason *string `json:"non_scan_reason,omitempty"`
}

func (o ImageLocalInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageLocalInfo struct{}"
	}

	return strings.Join([]string{"ImageLocalInfo", string(data)}, " ")
}
