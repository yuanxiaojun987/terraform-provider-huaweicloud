package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProcessStatisticsRequest Request Object
type ListProcessStatisticsRequest struct {

	// 路径
	Path *string `json:"path,omitempty"`

	// 企业项目
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 默认10
	Limit *int32 `json:"limit,omitempty"`

	// 默认是0
	Offset *int32 `json:"offset,omitempty"`

	// 类别，默认为host，包含如下： - host：主机 - container：容器
	Category *string `json:"category,omitempty"`
}

func (o ListProcessStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProcessStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ListProcessStatisticsRequest", string(data)}, " ")
}
