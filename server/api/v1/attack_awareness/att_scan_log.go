package attack_awareness

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/attack_awareness"
	attack_awarenessReq "github.com/flipped-aurora/gin-vue-admin/server/model/attack_awareness/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ScanlogApi struct {
}

var scanlogService = service.ServiceGroupApp.Attack_awarenessServiceGroup.ScanlogService

// CreateScanlog 创建扫描感知
// @Tags Scanlog
// @Summary 创建扫描感知
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body attack_awareness.Scanlog true "创建扫描感知"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /scanlog/createScanlog [post]
func (scanlogApi *ScanlogApi) CreateScanlog(c *gin.Context) {
	var scanlog attack_awareness.Scanlog
	err := c.ShouldBindJSON(&scanlog)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := scanlogService.CreateScanlog(&scanlog); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteScanlog 删除扫描感知
// @Tags Scanlog
// @Summary 删除扫描感知
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body attack_awareness.Scanlog true "删除扫描感知"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /scanlog/deleteScanlog [delete]
func (scanlogApi *ScanlogApi) DeleteScanlog(c *gin.Context) {
	ID := c.Query("ID")
	if err := scanlogService.DeleteScanlog(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteScanlogByIds 批量删除扫描感知
// @Tags Scanlog
// @Summary 批量删除扫描感知
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /scanlog/deleteScanlogByIds [delete]
func (scanlogApi *ScanlogApi) DeleteScanlogByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := scanlogService.DeleteScanlogByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateScanlog 更新扫描感知
// @Tags Scanlog
// @Summary 更新扫描感知
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body attack_awareness.Scanlog true "更新扫描感知"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /scanlog/updateScanlog [put]
func (scanlogApi *ScanlogApi) UpdateScanlog(c *gin.Context) {
	var scanlog attack_awareness.Scanlog
	err := c.ShouldBindJSON(&scanlog)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := scanlogService.UpdateScanlog(scanlog); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindScanlog 用id查询扫描感知
// @Tags Scanlog
// @Summary 用id查询扫描感知
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query attack_awareness.Scanlog true "用id查询扫描感知"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /scanlog/findScanlog [get]
func (scanlogApi *ScanlogApi) FindScanlog(c *gin.Context) {
	ID := c.Query("ID")
	if rescanlog, err := scanlogService.GetScanlog(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rescanlog": rescanlog}, c)
	}
}

// GetScanlogList 分页获取扫描感知列表
// @Tags Scanlog
// @Summary 分页获取扫描感知列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query attack_awarenessReq.ScanlogSearch true "分页获取扫描感知列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /scanlog/getScanlogList [get]
func (scanlogApi *ScanlogApi) GetScanlogList(c *gin.Context) {
	var pageInfo attack_awarenessReq.ScanlogSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := scanlogService.GetScanlogInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// GetScanlogPublic 不需要鉴权的扫描感知接口
// @Tags Scanlog
// @Summary 不需要鉴权的扫描感知接口
// @accept application/json
// @Produce application/json
// @Param data query attack_awarenessReq.ScanlogSearch true "分页获取扫描感知列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /scanlog/getScanlogPublic [get]
func (scanlogApi *ScanlogApi) GetScanlogPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的扫描感知接口信息",
	}, "获取成功", c)
}

// GetScanlogTop10 用id查询扫描感知Top10统计信息
// @Tags Scanlog
// @Summary 用id查询扫描感知Top10统计信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query attack_awareness.Scanlog true "用id查询扫描感知Top10统计信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /scanlog/getScanlogTop10 [get]
func (scanlogApi *ScanlogApi) GetScanlogTop10(c *gin.Context) {
	results, err := scanlogService.GetScanlogTop10()
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"results": results}, c)
	}
}
