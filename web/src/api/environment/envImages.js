import service from '@/utils/request'

// @Tags Images
// @Summary 创建镜像
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Images true "创建镜像"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /images/createImages [post]
export const createImages = (data) => {
  return service({
    url: '/images/createImages',
    method: 'post',
    data
  })
}

// @Tags Images
// @Summary 删除镜像
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Images true "删除镜像"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /images/deleteImages [delete]
export const deleteImages = (params) => {
  return service({
    url: '/images/deleteImages',
    method: 'delete',
    params
  })
}

// @Tags Images
// @Summary 批量删除镜像
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除镜像"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /images/deleteImages [delete]
export const deleteImagesByIds = (params) => {
  return service({
    url: '/images/deleteImagesByIds',
    method: 'delete',
    params
  })
}

// @Tags Images
// @Summary 更新镜像
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Images true "更新镜像"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /images/updateImages [put]
export const updateImages = (data) => {
  return service({
    url: '/images/updateImages',
    method: 'put',
    data
  })
}

// @Tags Images
// @Summary 用id查询镜像
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Images true "用id查询镜像"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /images/findImages [get]
export const findImages = (params) => {
  return service({
    url: '/images/findImages',
    method: 'get',
    params
  })
}

// @Tags Images
// @Summary 分页获取镜像列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取镜像列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /images/getImagesList [get]
export const getImagesList = (params) => {
  return service({
    url: '/images/getImagesList',
    method: 'get',
    params
  })
}

// @Tags Scanlog
// @Summary 获取扫描感知记录Top10统计信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "获取扫描感知记录Top10统计信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /scanlog/top10 [get]
export const getScanlogTop10 = (params) => {
  return service({
    url: '/scanlog/getScanlogTop10',
    method: 'get',
    params
  })
}