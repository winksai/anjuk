package order

import (
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/order"
	orderReq "github.com/flipped-aurora/gin-vue-admin/server/model/order/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RentalOrderApi struct{}

// CreateRentalOrder 创建rentalOrder表
// @Tags RentalOrder
// @Summary 创建rentalOrder表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body order.RentalOrder true "创建rentalOrder表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /rentalOrder/createRentalOrder [post]
func (rentalOrderApi *RentalOrderApi) CreateRentalOrder(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var rentalOrder order.RentalOrder
	err := c.ShouldBindJSON(&rentalOrder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = rentalOrderService.CreateRentalOrder(ctx, &rentalOrder)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteRentalOrder 删除rentalOrder表
// @Tags RentalOrder
// @Summary 删除rentalOrder表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body order.RentalOrder true "删除rentalOrder表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /rentalOrder/deleteRentalOrder [delete]
func (rentalOrderApi *RentalOrderApi) DeleteRentalOrder(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	err := rentalOrderService.DeleteRentalOrder(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteRentalOrderByIds 批量删除rentalOrder表
// @Tags RentalOrder
// @Summary 批量删除rentalOrder表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /rentalOrder/deleteRentalOrderByIds [delete]
func (rentalOrderApi *RentalOrderApi) DeleteRentalOrderByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := rentalOrderService.DeleteRentalOrderByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateRentalOrder 更新rentalOrder表
// @Tags RentalOrder
// @Summary 更新rentalOrder表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body order.RentalOrder true "更新rentalOrder表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /rentalOrder/updateRentalOrder [put]
func (rentalOrderApi *RentalOrderApi) UpdateRentalOrder(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var rentalOrder order.RentalOrder
	err := c.ShouldBindJSON(&rentalOrder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = rentalOrderService.UpdateRentalOrder(ctx, rentalOrder)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindRentalOrder 用id查询rentalOrder表
// @Tags RentalOrder
// @Summary 用id查询rentalOrder表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询rentalOrder表"
// @Success 200 {object} response.Response{data=order.RentalOrder,msg=string} "查询成功"
// @Router /rentalOrder/findRentalOrder [get]
func (rentalOrderApi *RentalOrderApi) FindRentalOrder(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	detail, err := rentalOrderService.GetRentalOrderDetail(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(detail, c)
}

// GetRentalOrderList 分页获取rentalOrder表列表
// @Tags RentalOrder
// @Summary 分页获取rentalOrder表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query orderReq.RentalOrderSearch true "分页获取rentalOrder表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /rentalOrder/getRentalOrderList [get]
func (rentalOrderApi *RentalOrderApi) GetRentalOrderList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo orderReq.RentalOrderSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := rentalOrderService.GetRentalOrderInfoList(ctx, pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetRentalOrderPublic 不需要鉴权的rentalOrder表接口
// @Tags RentalOrder
// @Summary 不需要鉴权的rentalOrder表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /rentalOrder/getRentalOrderPublic [get]
func (rentalOrderApi *RentalOrderApi) GetRentalOrderPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	rentalOrderService.GetRentalOrderPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的rentalOrder表接口信息",
	}, "获取成功", c)
}

// 下载合同PDF
func DownloadContract(c *gin.Context) {
	orderId := c.Query("id")
	if orderId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "缺少订单ID"})
		return
	}

	// 查询合同表，获取合同 OSS 地址
	var contract order.Contract
	if err := global.GVA_DB.Where("order_id = ?", orderId).First(&contract).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "未找到该订单的合同"})
		return
	}
	contractUrl := contract.ContractURL
	if contractUrl == "" {
		c.JSON(http.StatusNotFound, gin.H{"msg": "未找到合同文件"})
		return
	}

	// 下载 OSS 文件流并转发
	resp, err := http.Get(contractUrl)
	if err != nil || resp.StatusCode != 200 {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "下载合同失败"})
		return
	}
	defer resp.Body.Close()

	// 取文件名
	u, _ := url.Parse(contractUrl)
	parts := strings.Split(u.Path, "/")
	filename := parts[len(parts)-1]
	if !strings.HasSuffix(strings.ToLower(filename), ".pdf") {
		filename += ".pdf"
	}

	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Status(http.StatusOK)
	io.Copy(c.Writer, resp.Body)
}
