package order

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ RentalOrderRouter }

var rentalOrderApi = api.ApiGroupApp.OrderApiGroup.RentalOrderApi
