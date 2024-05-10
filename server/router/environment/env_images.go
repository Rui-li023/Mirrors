package environment

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ImagesRouter struct {
}

// InitImagesRouter 初始化 镜像 路由信息
func (s *ImagesRouter) InitImagesRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	imagesRouter := Router.Group("images").Use(middleware.OperationRecord())
	imagesRouterWithoutRecord := Router.Group("images")
	imagesRouterWithoutAuth := PublicRouter.Group("images")

	var imagesApi = v1.ApiGroupApp.EnvironmentApiGroup.ImagesApi
	{
		imagesRouter.POST("createImages", imagesApi.CreateImages)   // 新建镜像
		imagesRouter.DELETE("deleteImages", imagesApi.DeleteImages) // 删除镜像
		imagesRouter.DELETE("deleteImagesByIds", imagesApi.DeleteImagesByIds) // 批量删除镜像
		imagesRouter.PUT("updateImages", imagesApi.UpdateImages)    // 更新镜像
	}
	{
		imagesRouterWithoutRecord.GET("findImages", imagesApi.FindImages)        // 根据ID获取镜像
		imagesRouterWithoutRecord.GET("getImagesList", imagesApi.GetImagesList)  // 获取镜像列表
	}
	{
	    imagesRouterWithoutAuth.GET("getImagesPublic", imagesApi.GetImagesPublic)  // 获取镜像列表
	}
}
