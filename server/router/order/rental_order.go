package order

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/order"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type RentalOrderRouter struct{}

// InitRentalOrderRouter 初始化 rentalOrder表 路由信息
func (s *RentalOrderRouter) InitRentalOrderRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	rentalOrderRouter := Router.Group("rentalOrder").Use(middleware.OperationRecord())
	rentalOrderRouterWithoutRecord := Router.Group("rentalOrder")
	rentalOrderRouterWithoutAuth := PublicRouter.Group("rentalOrder")
	{
		rentalOrderRouter.POST("createRentalOrder", rentalOrderApi.CreateRentalOrder)             // 新建rentalOrder表
		rentalOrderRouter.DELETE("deleteRentalOrder", rentalOrderApi.DeleteRentalOrder)           // 删除rentalOrder表
		rentalOrderRouter.DELETE("deleteRentalOrderByIds", rentalOrderApi.DeleteRentalOrderByIds) // 批量删除rentalOrder表
		rentalOrderRouter.PUT("updateRentalOrder", rentalOrderApi.UpdateRentalOrder)              // 更新rentalOrder表
		rentalOrderRouter.GET("downloadContract", order.DownloadContract)                         //下载合同
	}
	{
		rentalOrderRouterWithoutRecord.GET("findRentalOrder", rentalOrderApi.FindRentalOrder)       // 根据ID获取rentalOrder表
		rentalOrderRouterWithoutRecord.GET("getRentalOrderList", rentalOrderApi.GetRentalOrderList) // 获取rentalOrder表列表
	}
	{
		rentalOrderRouterWithoutAuth.GET("getRentalOrderPublic", rentalOrderApi.GetRentalOrderPublic) // rentalOrder表开放接口
	}
}
