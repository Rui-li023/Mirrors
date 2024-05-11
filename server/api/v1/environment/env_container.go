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
	"strings"
)

type ContainerApi struct {
}

var containerService = service.ServiceGroupApp.EnvironmentServiceGroup.ContainerService

// CreateContainer 创建容器管理
// @Tags Container
// @Summary 创建容器管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body environment.Container true "创建容器管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /container/createContainer [post]
func (containerApi *ContainerApi) CreateContainer(c *gin.Context) {
	var container environment.Container
	err := c.ShouldBindJSON(&container)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := containerService.CreateContainer(&container); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteContainer 删除容器管理
// @Tags Container
// @Summary 删除容器管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body environment.Container true "删除容器管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /container/deleteContainer [delete]
func (containerApi *ContainerApi) DeleteContainer(c *gin.Context) {
	ID := c.Query("ID")
	if err := containerService.DeleteContainer(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteContainerByIds 批量删除容器管理
// @Tags Container
// @Summary 批量删除容器管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /container/deleteContainerByIds [delete]
func (containerApi *ContainerApi) DeleteContainerByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := containerService.DeleteContainerByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateContainer 更新容器管理
// @Tags Container
// @Summary 更新容器管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body environment.Container true "更新容器管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /container/updateContainer [put]
func (containerApi *ContainerApi) UpdateContainer(c *gin.Context) {
	var container environment.Container
	err := c.ShouldBindJSON(&container)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := containerService.UpdateContainer(container); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindContainer 用id查询容器管理
// @Tags Container
// @Summary 用id查询容器管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query environment.Container true "用id查询容器管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /container/findContainer [get]
func (containerApi *ContainerApi) FindContainer(c *gin.Context) {
	ID := c.Query("ID")
	if recontainer, err := containerService.GetContainer(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recontainer": recontainer}, c)
	}
}

// GetContainerList 分页获取容器管理列表
// @Tags Container
// @Summary 分页获取容器管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query environmentReq.ContainerSearch true "分页获取容器管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /container/getContainerList [get]
func (containerApi *ContainerApi) GetContainerList(c *gin.Context) {
	var pageInfo environmentReq.ContainerSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	containers, err := utils.GetDockerContainers()
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	filteredContainers := make([]environment.Container, 0)
	for _, container := range containers {
		if pageInfo.ContainerId != "" && container.ContainerId != pageInfo.ContainerId {
			continue
		}
		if pageInfo.ImageName != "" && !strings.Contains(container.ImageName, pageInfo.ImageName) {
			continue
		}
		filteredContainers = append(filteredContainers, container)
	}

	var total = int64(len(filteredContainers))
	var pageSize = pageInfo.PageSize
	var page = pageInfo.Page
	var start = (page - 1) * pageSize
	var end = start + pageSize
	if end > len(filteredContainers) {
		end = len(filteredContainers)
	}
	var paginatedContainers = filteredContainers[start:end]

	response.OkWithDetailed(response.PageResult{
		List:     paginatedContainers,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}, "获取成功", c)
}

// GetContainerPublic 不需要鉴权的容器管理接口
// @Tags Container
// @Summary 不需要鉴权的容器管理接口
// @accept application/json
// @Produce application/json
// @Param data query environmentReq.ContainerSearch true "分页获取容器管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /container/getContainerPublic [get]
func (containerApi *ContainerApi) GetContainerPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的容器管理接口信息",
	}, "获取成功", c)
}

// StartContainer 根据id启动容器
// @Tags Container
// @Summary 根据id启动容器
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query environment.Container true "根据id启动容器"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"启动成功"}"
// @Router /container/startContainer [get]
func (containerApi *ContainerApi) StartContainer(c *gin.Context) {
	id := c.Query("ID")

	err := utils.StartContainerByID(id)
	if err != nil {
		global.GVA_LOG.Error("启动失败!", zap.Error(err))
		response.FailWithMessage("启动失败", c)
		return
	} else {
		response.OkWithMessage("启动成功", c)
	}
}

// StopContainer 根据id停止容器
// @Tags Container
// @Summary 根据id停止容器
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query environment.Container true "根据id停止容器"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"停止成功"}"
// @Router /container/stopContainer [get]
func (containerApi *ContainerApi) StopContainer(c *gin.Context) {
	id := c.Query("ID")

	err := utils.StopContainerByID(id)
	if err != nil {
		global.GVA_LOG.Error("停止失败!", zap.Error(err))
		response.FailWithMessage("停止失败", c)
		return
	} else {
		response.OkWithMessage("停止成功", c)
	}
}

// RestartContainer 根据id重启容器
// @Tags Container
// @Summary 根据id重启容器
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query environment.Container true "根据id重启容器"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"重启成功"}"
// @Router /container/restartContainer [get]
func (containerApi *ContainerApi) RestartContainer(c *gin.Context) {
	id := c.Query("ID")

	err := utils.RestartContainerByID(id)
	if err != nil {
		global.GVA_LOG.Error("重启失败!", zap.Error(err))
		response.FailWithMessage("重启失败", c)
		return
	} else {
		response.OkWithMessage("重启成功", c)
	}
}
