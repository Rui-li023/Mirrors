package attack_awareness

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ScanlogRouter struct {
}

// InitScanlogRouter 初始化 扫描感知 路由信息
func (s *ScanlogRouter) InitScanlogRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	scanlogRouter := Router.Group("scanlog").Use(middleware.OperationRecord())
	scanlogRouterWithoutRecord := Router.Group("scanlog")
	scanlogRouterWithoutAuth := PublicRouter.Group("scanlog")

	var scanlogApi = v1.ApiGroupApp.Attack_awarenessApiGroup.ScanlogApi
	{
		//scanlogRouter.POST("createScanlog", scanlogApi.CreateScanlog)             // 新建扫描感知
		scanlogRouter.DELETE("deleteScanlog", scanlogApi.DeleteScanlog)           // 删除扫描感知
		scanlogRouter.DELETE("deleteScanlogByIds", scanlogApi.DeleteScanlogByIds) // 批量删除扫描感知
		scanlogRouter.PUT("updateScanlog", scanlogApi.UpdateScanlog)              // 更新扫描感知
	}
	{
		scanlogRouterWithoutRecord.GET("findScanlog", scanlogApi.FindScanlog)       // 根据ID获取扫描感知
		scanlogRouterWithoutRecord.GET("getScanlogList", scanlogApi.GetScanlogList) // 获取扫描感知列表
	}
	{
		scanlogRouterWithoutAuth.GET("getScanlogPublic", scanlogApi.GetScanlogPublic) // 获取扫描感知列表
		scanlogRouterWithoutAuth.POST("createScanlog", scanlogApi.CreateScanlog)
	}
}
