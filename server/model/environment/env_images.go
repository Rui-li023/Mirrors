// 自动生成模板Images
package environment

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 镜像 结构体  Images
type Images struct {
	global.GVA_MODEL
	Image_Id     string `json:"image_Id" form:"image_Id" gorm:"primarykey;column:image_id;comment:;size:64;"` //镜像ID
	Repository   string `json:"repository" form:"repository" gorm:"column:repository;comment:;size:64;"`      //仓库
	Tag          string `json:"tag" form:"tag" gorm:"column:tag;comment:;size:128;"`                          //tag
	Size         string `json:"size" form:"size" gorm:"column:size;comment:;size:64;"`
	CreatedSince string `json:"created_since" form:"created_since" gorm:"column:created_since;comment:;size:64;"`
}

// TableName 镜像 Images自定义表名 images
func (Images) TableName() string {
	return "images"
}
