// 自动生成模板Container
package environment

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 容器管理 结构体  Container
type Container struct {
	global.GVA_MODEL
	ContainerId string `json:"containerId" form:"containerId" gorm:"column:container_id;comment:;size:128;"` //容器ID
	ImageName   string `json:"imageName" form:"imageName" gorm:"column:image_name;comment:;size:64;"`        //镜像名称
	Command     string `json:"command" form:"command" gorm:"column:command;comment:;size:256;"`              //启动命令
	Create      string `json:"create" form:"create" gorm:"column:create;comment:;size:256;"`                 //创建时间
	Status      string `json:"status" form:"status" gorm:"column:status;comment:;size:256;"`                 //运行状态
	Ports       string `json:"ports" form:"ports" gorm:"column:ports;comment:;size:256;"`                    //端口
}

// TableName 容器管理 Container自定义表名 container
func (Container) TableName() string {
	return "container"
}
