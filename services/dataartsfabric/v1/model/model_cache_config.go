package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CacheConfig **参数解释**：分布式缓存配置。 **约束限制**：不涉及。
type CacheConfig struct {

	// **参数解释**：挂载路径。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	MntPath string `json:"mnt_path"`

	// **参数解释**：分布式缓存id。 **约束限制**：不涉及。 **取值范围**：已创建的分布式缓存资源id。 **默认取值**：不涉及。
	CacheInsId string `json:"cache_ins_id"`
}

func (o CacheConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CacheConfig struct{}"
	}

	return strings.Join([]string{"CacheConfig", string(data)}, " ")
}
