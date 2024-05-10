package environment

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ContainerRouter struct {
}

// InitContainerRouter 初始化 容器管理 路由信息
func (s *ContainerRouter) InitContainerRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	containerRouter := Router.Group("container").Use(middleware.OperationRecord())
	containerRouterWithoutRecord := Router.Group("container")
	containerRouterWithoutAuth := PublicRouter.Group("container")

	var containerApi = v1.ApiGroupApp.EnvironmentApiGroup.ContainerApi
	{
		containerRouter.POST("createContainer", containerApi.CreateContainer)   // 新建容器管理
		containerRouter.DELETE("deleteContainer", containerApi.DeleteContainer) // 删除容器管理
		containerRouter.DELETE("deleteContainerByIds", containerApi.DeleteContainerByIds) // 批量删除容器管理
		containerRouter.PUT("updateContainer", containerApi.UpdateContainer)    // 更新容器管理
	}
	{
		containerRouterWithoutRecord.GET("findContainer", containerApi.FindContainer)        // 根据ID获取容器管理
		containerRouterWithoutRecord.GET("getContainerList", containerApi.GetContainerList)  // 获取容器管理列表
	}
	{
	    containerRouterWithoutAuth.GET("getContainerPublic", containerApi.GetContainerPublic)  // 获取容器管理列表
	}
}
