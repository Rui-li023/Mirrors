package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type ContainerSearch struct {
	ImageName   string `json:"imageName" form:"imageName" `
	ContainerId string `json:"containerId" form:"containerId" `
	request.PageInfo
}
