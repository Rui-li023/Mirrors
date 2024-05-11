package environment

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/environment"
	environmentReq "github.com/flipped-aurora/gin-vue-admin/server/model/environment/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ImagesApi struct {
}

var imagesService = service.ServiceGroupApp.EnvironmentServiceGroup.ImagesService

// CreateImages 创建镜像
// @Tags Images
// @Summary 创建镜像
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body environment.Images true "创建镜像"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /images/createImages [post]
func (imagesApi *ImagesApi) CreateImages(c *gin.Context) {
	var images environment.Images
	err := c.ShouldBindJSON(&images)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := utils.PullImage(images.Repository, images.Tag); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功，镜像加载需要时间，稍后请手动刷新！", c)
	}
}

// DeleteImages 删除镜像
// @Tags Images
// @Summary 删除镜像
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body environment.Images true "删除镜像"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /images/deleteImages [delete]
func (imagesApi *ImagesApi) DeleteImages(c *gin.Context) {
	ID := c.Query("ID")
	if err := utils.DeleteImageByID(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败，如果镜像被使用则无法删除！", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteImagesByIds 批量删除镜像
// @Tags Images
// @Summary 批量删除镜像
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /images/deleteImagesByIds [delete]
func (imagesApi *ImagesApi) DeleteImagesByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := imagesService.DeleteImagesByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateImages 更新镜像
// @Tags Images
// @Summary 更新镜像
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body environment.Images true "更新镜像"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /images/updateImages [put]
func (imagesApi *ImagesApi) UpdateImages(c *gin.Context) {
	var images environment.Images
	err := c.ShouldBindJSON(&images)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := imagesService.UpdateImages(images); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindImages 用id查询镜像
// @Tags Images
// @Summary 用id查询镜像
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query environment.Images true "用id查询镜像"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /images/findImages [get]
func (imagesApi *ImagesApi) FindImages(c *gin.Context) {
	ID := c.Query("ID")
	if reimages, err := imagesService.GetImages(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reimages": reimages}, c)
	}
}

// GetImagesList 分页获取镜像列表
// @Tags Images
// @Summary 分页获取镜像列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query environmentReq.ImagesSearch true "分页获取镜像列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /images/getImagesList [get]
func (imagesApi *ImagesApi) GetImagesList(c *gin.Context) {
	var pageInfo environmentReq.ImagesSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if list, err := utils.GetDockerImages(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    int64(len(list)),
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}

	//if list, total, err := imagesService.GetImagesInfoList(pageInfo); err != nil {
	//	global.GVA_LOG.Error("获取失败!", zap.Error(err))
	//	response.FailWithMessage("获取失败", c)
	//} else {
	//	response.OkWithDetailed(response.PageResult{
	//		List:     list,
	//		Total:    total,
	//		Page:     pageInfo.Page,
	//		PageSize: pageInfo.PageSize,
	//	}, "获取成功", c)
	//}
}

// GetImagesPublic 不需要鉴权的镜像接口
// @Tags Images
// @Summary 不需要鉴权的镜像接口
// @accept application/json
// @Produce application/json
// @Param data query environmentReq.ImagesSearch true "分页获取镜像列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /images/getImagesPublic [get]
func (imagesApi *ImagesApi) GetImagesPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的镜像接口信息",
	}, "获取成功", c)
}
