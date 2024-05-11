import service from '@/utils/request'

// @Tags Container
// @Summary 创建容器管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Container true "创建容器管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /container/createContainer [post]
export const createContainer = (data) => {
  return service({
    url: '/container/createContainer',
    method: 'post',
    data
  })
}

// @Tags Container
// @Summary 删除容器管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Container true "删除容器管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /container/deleteContainer [delete]
export const deleteContainer = (params) => {
  return service({
    url: '/container/deleteContainer',
    method: 'delete',
    params
  })
}

// @Tags Container
// @Summary 根据id启动容器
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Container true "根据id启动容器"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"启动成功"}"
// @Router /container/startContainer [get]
export const startContainer = (params) =>{
  return service({
    url: '/container/startContainer',
    method: 'get',
    params
  })
}

// @Tags Container
// @Summary 根据id停止容器
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Container true "根据id停止容器"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"停止成功"}"
// @Router /container/stopContainer [get]
export const stopContainer = (params) =>{
  return service({
    url: '/container/stopContainer',
    method: 'get',
    params
  })
}

// @Tags Container
// @Summary 根据id重启容器
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Container true "根据id重启容器"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"重启成功"}"
// @Router /container/restartContainer [get]
export const restartContainer = (params) =>{
  return service({
    url: '/container/restartContainer',
    method: 'get',
    params
  })
}

// @Tags Container
// @Summary 批量删除容器管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除容器管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /container/deleteContainer [delete]
export const deleteContainerByIds = (params) => {
  return service({
    url: '/container/deleteContainerByIds',
    method: 'delete',
    params
  })
}

// @Tags Container
// @Summary 更新容器管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Container true "更新容器管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /container/updateContainer [put]
export const updateContainer = (data) => {
  return service({
    url: '/container/updateContainer',
    method: 'put',
    data
  })
}

// @Tags Container
// @Summary 用id查询容器管理
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Container true "用id查询容器管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /container/findContainer [get]
export const findContainer = (params) => {
  return service({
    url: '/container/findContainer',
    method: 'get',
    params
  })
}

// @Tags Container
// @Summary 分页获取容器管理列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取容器管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /container/getContainerList [get]
export const getContainerList = (params) => {
  return service({
    url: '/container/getContainerList',
    method: 'get',
    params
  })
}
