package api

import (
	"app/global"
	"app/model/discuss"
	"app/mysqlDB"
	"app/response"
	"app/service"
	"github.com/gin-gonic/gin"
)

type DisInfoApi struct {
}

var disDataService = service.ServiceApi.DisService

// CreateDisInfo 创建帖子信息
// @Tags DisInfo
// @Summary 创建帖子信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body discuss.DisInfo true "创建帖子信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /disData/createDisInfo [post]
func (disDataApi *DisInfoApi) CreateDisInfo(c *gin.Context) {
	var disData discuss.DisInfo
	err := c.ShouldBindJSON(&disData)
	userid, ok := c.Get("userID")
	if err != nil {
		response.FailWithDetailed(err.Error(), "创建失败", c)
		return
	}
	if !ok {
		response.FailWithMessage("身份验证失败", c)
		return
	}
	disData.DisUserId = userid.(uint)

	if disData.DisModel == 1 {
		user, err := mysqlDB.FindUserByID(disData.DisUserId)
		if err != nil {
			response.FailWithMessage("查找用户失败", c)
			return
		}
		if user.UserModel != global.UserModelGuanLiYuan {
			response.FailWithMessage("用户无法发送公告帖子", c)
			return
		}
	}
	if err := disDataService.CreateDisInfo(&disData); err != nil {
		//global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteDisInfo 删除帖子信息
// @Tags DisInfo
// @Summary 删除帖子信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body discuss.DisInfo true "删除帖子信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /disData/deleteDisInfo [delete]
func (disDataApi *DisInfoApi) DeleteDisInfo(c *gin.Context) {
	id := c.Query("ID")
	userid, ok := c.Get("userID")
	dis, err := mysqlDB.FindDisByID(id)
	if err != nil {
		response.FailWithDetailed(err.Error(), "删除的帖子不存在", c)
		return
	}
	disuser, err := mysqlDB.FindUserByID(dis.DisUserId)
	if err != nil {
		response.FailWithMessage("删除权限不足", c)
		return
	}
	if ok == false || ((userid.(uint)) != dis.DisUserId && disuser.UserModel != global.UserModelGuanLiYuan) {
		response.FailWithMessage("身份验证失败", c)
		return
	}
	if err := disDataService.DeleteDisInfo(id); err != nil {
		//global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteDisInfoByIds 批量删除帖子信息
// @Tags DisInfo
// @Summary 批量删除帖子信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除帖子信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /disData/deleteDisInfoByIds [delete]
func (disDataApi *DisInfoApi) DeleteDisInfoByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	if err := disDataService.DeleteDisInfoByIds(ids); err != nil {
		//global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateDisInfo 更新帖子信息
// @Tags DisInfo
// @Summary 更新帖子信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body discuss.DisInfo true "更新帖子信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /disData/updateDisInfo [put]
func (disDataApi *DisInfoApi) UpdateDisInfo(c *gin.Context) {
	var disData discuss.DisInfo
	err := c.ShouldBindJSON(&disData)
	userid, ok := c.Get("userID")
	if !ok || (disData.DisUserId != 0 && (userid.(uint)) != disData.DisUserId) {
		response.FailWithMessage("身份验证失败", c)
		return
	}
	disData.DisUserId = userid.(uint)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := disDataService.UpdateDisInfo(disData); err != nil {
		//global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindDisInfo 用id查询帖子信息
// @Tags DisInfo
// @Summary 用id查询帖子信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query discuss.DisInfo true "用id查询帖子信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /disData/findDisInfo [get]
func (disDataApi *DisInfoApi) FindDisInfo(c *gin.Context) {
	id := c.Query("ID")
	if redisData, err := disDataService.GetDisInfo(id); err != nil {
		//global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redisData": redisData}, c)
	}
}

// GetDisInfoList 分页获取帖子信息列表
// @Tags DisInfo
// @Summary 分页获取帖子信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query discussReq.DisInfoSearch true "分页获取帖子信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /disData/getDisInfoList [get]
func (disDataApi *DisInfoApi) GetDisInfoList(c *gin.Context) {
	var pageInfo discuss.DisInfoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := disDataService.GetDisInfoInfoList(pageInfo); err != nil {
		//global.GVA_LOG.Error("获取失败!", zap.Error(err))
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
