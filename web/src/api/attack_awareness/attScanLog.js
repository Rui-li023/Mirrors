import service from '@/utils/request'

// @Tags Scanlog
// @Summary 创建扫描感知
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Scanlog true "创建扫描感知"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /scanlog/createScanlog [post]
export const createScanlog = (data) => {
  return service({
    url: '/scanlog/createScanlog',
    method: 'post',
    data
  })
}

// @Tags Scanlog
// @Summary 删除扫描感知
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Scanlog true "删除扫描感知"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /scanlog/deleteScanlog [delete]
export const deleteScanlog = (params) => {
  return service({
    url: '/scanlog/deleteScanlog',
    method: 'delete',
    params
  })
}

// @Tags Scanlog
// @Summary 批量删除扫描感知
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除扫描感知"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /scanlog/deleteScanlog [delete]
export const deleteScanlogByIds = (params) => {
  return service({
    url: '/scanlog/deleteScanlogByIds',
    method: 'delete',
    params
  })
}

// @Tags Scanlog
// @Summary 更新扫描感知
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Scanlog true "更新扫描感知"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /scanlog/updateScanlog [put]
export const updateScanlog = (data) => {
  return service({
    url: '/scanlog/updateScanlog',
    method: 'put',
    data
  })
}

// @Tags Scanlog
// @Summary 用id查询扫描感知
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Scanlog true "用id查询扫描感知"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /scanlog/findScanlog [get]
export const findScanlog = (params) => {
  return service({
    url: '/scanlog/findScanlog',
    method: 'get',
    params
  })
}

// @Tags Scanlog
// @Summary 分页获取扫描感知列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取扫描感知列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /scanlog/getScanlogList [get]
export const getScanlogList = (params) => {
  return service({
    url: '/scanlog/getScanlogList',
    method: 'get',
    params
  })
}
