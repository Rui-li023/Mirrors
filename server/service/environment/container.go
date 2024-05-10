package environment

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/environment"
    environmentReq "github.com/flipped-aurora/gin-vue-admin/server/model/environment/request"
)

type ContainerService struct {
}

// CreateContainer 创建容器管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (containerService *ContainerService) CreateContainer(container *environment.Container) (err error) {
	err = global.GVA_DB.Create(container).Error
	return err
}

// DeleteContainer 删除容器管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (containerService *ContainerService)DeleteContainer(ID string) (err error) {
	err = global.GVA_DB.Delete(&environment.Container{},"id = ?",ID).Error
	return err
}

// DeleteContainerByIds 批量删除容器管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (containerService *ContainerService)DeleteContainerByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]environment.Container{},"id in ?",IDs).Error
	return err
}

// UpdateContainer 更新容器管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (containerService *ContainerService)UpdateContainer(container environment.Container) (err error) {
	err = global.GVA_DB.Model(&environment.Container{}).Where("id = ?",container.ID).Updates(&container).Error
	return err
}

// GetContainer 根据ID获取容器管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (containerService *ContainerService)GetContainer(ID string) (container environment.Container, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&container).Error
	return
}

// GetContainerInfoList 分页获取容器管理记录
// Author [piexlmax](https://github.com/piexlmax)
func (containerService *ContainerService)GetContainerInfoList(info environmentReq.ContainerSearch) (list []environment.Container, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&environment.Container{})
    var containers []environment.Container
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.ImageName != "" {
        db = db.Where("image_name LIKE ?","%"+ info.ImageName+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&containers).Error
	return  containers, total, err
}