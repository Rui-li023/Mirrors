package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ImagesSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	Image_Id       string     `json:"image_Id" form:"image_Id" `
	Repository     string     `json:"repository" form:"repository" `
	request.PageInfo
}
