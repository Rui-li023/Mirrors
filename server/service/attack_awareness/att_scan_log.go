package attack_awareness

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/attack_awareness"
    attack_awarenessReq "github.com/flipped-aurora/gin-vue-admin/server/model/attack_awareness/request"
)

type ScanlogService struct {
}

// CreateScanlog 创建扫描感知记录
// Author [piexlmax](https://github.com/piexlmax)
func (scanlogService *ScanlogService) CreateScanlog(scanlog *attack_awareness.Scanlog) (err error) {
	err = global.GVA_DB.Create(scanlog).Error
	return err
}

// DeleteScanlog 删除扫描感知记录
// Author [piexlmax](https://github.com/piexlmax)
func (scanlogService *ScanlogService)DeleteScanlog(ID string) (err error) {
	err = global.GVA_DB.Delete(&attack_awareness.Scanlog{},"id = ?",ID).Error
	return err
}

// DeleteScanlogByIds 批量删除扫描感知记录
// Author [piexlmax](https://github.com/piexlmax)
func (scanlogService *ScanlogService)DeleteScanlogByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]attack_awareness.Scanlog{},"id in ?",IDs).Error
	return err
}

// UpdateScanlog 更新扫描感知记录
// Author [piexlmax](https://github.com/piexlmax)
func (scanlogService *ScanlogService)UpdateScanlog(scanlog attack_awareness.Scanlog) (err error) {
	err = global.GVA_DB.Model(&attack_awareness.Scanlog{}).Where("id = ?",scanlog.ID).Updates(&scanlog).Error
	return err
}

// GetScanlog 根据ID获取扫描感知记录
// Author [piexlmax](https://github.com/piexlmax)
func (scanlogService *ScanlogService)GetScanlog(ID string) (scanlog attack_awareness.Scanlog, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&scanlog).Error
	return
}

// GetScanlogInfoList 分页获取扫描感知记录
// Author [piexlmax](https://github.com/piexlmax)
func (scanlogService *ScanlogService)GetScanlogInfoList(info attack_awarenessReq.ScanlogSearch) (list []attack_awareness.Scanlog, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&attack_awareness.Scanlog{})
    var scanlogs []attack_awareness.Scanlog
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Source_ip != "" {
        db = db.Where("source_ip = ?",info.Source_ip)
    }
    if info.Source_port != nil {
        db = db.Where("source_port <> ?",info.Source_port)
    }
    if info.Protocol != "" {
        db = db.Where("protocol = ?",info.Protocol)
    }
    if info.Dest_ip != "" {
        db = db.Where("dest_ip <> ?",info.Dest_ip)
    }
    if info.Dest_port != nil {
        db = db.Where("dest_port <> ?",info.Dest_port)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&scanlogs).Error
	return  scanlogs, total, err
}