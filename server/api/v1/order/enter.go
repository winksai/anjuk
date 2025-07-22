package order

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ RentalOrderApi }

var rentalOrderService = service.ServiceGroupApp.OrderServiceGroup.RentalOrderService
