import service from '@/utils/request'
// @Tags RentalOrder
// @Summary 创建rentalOrder表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.RentalOrder true "创建rentalOrder表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /rentalOrder/createRentalOrder [post]
export const createRentalOrder = (data) => {
  return service({
    url: '/rentalOrder/createRentalOrder',
    method: 'post',
    data
  })
}

// @Tags RentalOrder
// @Summary 删除rentalOrder表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.RentalOrder true "删除rentalOrder表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /rentalOrder/deleteRentalOrder [delete]
export const deleteRentalOrder = (params) => {
  return service({
    url: '/rentalOrder/deleteRentalOrder',
    method: 'delete',
    params
  })
}

// @Tags RentalOrder
// @Summary 批量删除rentalOrder表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除rentalOrder表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /rentalOrder/deleteRentalOrder [delete]
export const deleteRentalOrderByIds = (params) => {
  return service({
    url: '/rentalOrder/deleteRentalOrderByIds',
    method: 'delete',
    params
  })
}

// @Tags RentalOrder
// @Summary 更新rentalOrder表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.RentalOrder true "更新rentalOrder表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /rentalOrder/updateRentalOrder [put]
export const updateRentalOrder = (data) => {
  return service({
    url: '/rentalOrder/updateRentalOrder',
    method: 'put',
    data
  })
}

// @Tags RentalOrder
// @Summary 用id查询rentalOrder表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.RentalOrder true "用id查询rentalOrder表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /rentalOrder/findRentalOrder [get]
export const findRentalOrder = (params) => {
  return service({
    url: '/rentalOrder/findRentalOrder',
    method: 'get',
    params
  })
}

// @Tags RentalOrder
// @Summary 分页获取rentalOrder表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取rentalOrder表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /rentalOrder/getRentalOrderList [get]
export const getRentalOrderList = (params) => {
  return service({
    url: '/rentalOrder/getRentalOrderList',
    method: 'get',
    params
  })
}

// @Tags RentalOrder
// @Summary 不需要鉴权的rentalOrder表接口
// @Accept application/json
// @Produce application/json
// @Param data query orderReq.RentalOrderSearch true "分页获取rentalOrder表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /rentalOrder/getRentalOrderPublic [get]
export const getRentalOrderPublic = () => {
  return service({
    url: '/rentalOrder/getRentalOrderPublic',
    method: 'get',
  })
}

// 下载合同PDF
export const downloadContract = (orderId) => {
  return service({
    url: '/rentalOrder/downloadContract',
    method: 'get',
    params: { id: orderId },
    responseType: 'blob'
  })
}
