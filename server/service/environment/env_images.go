package environment

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/environment"
    environmentReq "github.com/flipped-aurora/gin-vue-admin/server/model/environment/request"
)

type ImagesService struct {
}

// CreateImages 创建镜像记录
// Author [piexlmax](https://github.com/piexlmax)
func (imagesService *ImagesService) CreateImages(images *environment.Images) (err error) {
	err = global.GVA_DB.Create(images).Error
	return err
}

// DeleteImages 删除镜像记录
// Author [piexlmax](https://github.com/piexlmax)
func (imagesService *ImagesService)DeleteImages(ID string) (err error) {
	err = global.GVA_DB.Delete(&environment.Images{},"id = ?",ID).Error
	return err
}

// DeleteImagesByIds 批量删除镜像记录
// Author [piexlmax](https://github.com/piexlmax)
func (imagesService *ImagesService)DeleteImagesByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]environment.Images{},"id in ?",IDs).Error
	return err
}

// UpdateImages 更新镜像记录
// Author [piexlmax](https://github.com/piexlmax)
func (imagesService *ImagesService)UpdateImages(images environment.Images) (err error) {
	err = global.GVA_DB.Model(&environment.Images{}).Where("id = ?",images.ID).Updates(&images).Error
	return err
}

// GetImages 根据ID获取镜像记录
// Author [piexlmax](https://github.com/piexlmax)
func (imagesService *ImagesService)GetImages(ID string) (images environment.Images, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&images).Error
	return
}

// GetImagesInfoList 分页获取镜像记录
// Author [piexlmax](https://github.com/piexlmax)
func (imagesService *ImagesService)GetImagesInfoList(info environmentReq.ImagesSearch) (list []environment.Images, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&environment.Images{})
    var imagess []environment.Images
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Image_Id != "" {
        db = db.Where("image__id LIKE ?","%"+ info.Image_Id+"%")
    }
    if info.Repository != "" {
        db = db.Where("repository LIKE ?","%"+ info.Repository+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&imagess).Error
	return  imagess, total, err
}