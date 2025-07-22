package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type RentalOrderSearch struct {
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	TenantId       *int        `json:"tenantId" form:"tenantId"`
	LandlordId     *int        `json:"landlordId" form:"landlordId"`
	HouseTitle     string      `json:"houseTitle" form:"houseTitle"`
	TenantName     string      `json:"tenantName" form:"tenantName"`
	LandlordName   string      `json:"landlordName" form:"landlordName"`
	Status         string      `json:"status" form:"status"`
	request.PageInfo
}
