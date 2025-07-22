// 自动生成模板RentalOrder
package order

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// rentalOrder表 结构体  RentalOrder
type RentalOrder struct {
	global.GVA_MODEL
	HouseId    *int       `json:"houseId" form:"houseId" gorm:"comment:房源ID;column:house_id;size:19;"`            //房源ID
	TenantId   *int       `json:"tenantId" form:"tenantId" gorm:"comment:租客ID;column:tenant_id;size:19;"`         //租客ID
	LandlordId *int       `json:"landlordId" form:"landlordId" gorm:"comment:房东ID;column:landlord_id;size:19;"`   //房东ID
	RentStart  *time.Time `json:"rentStart" form:"rentStart" gorm:"comment:租期开始;column:rent_start;"`              //租期开始
	RentEnd    *time.Time `json:"rentEnd" form:"rentEnd" gorm:"comment:租期结束;column:rent_end;"`                    //租期结束
	RentAmount *float64   `json:"rentAmount" form:"rentAmount" gorm:"comment:租金;column:rent_amount;size:10;"`     //租金
	Deposit    *float64   `json:"deposit" form:"deposit" gorm:"comment:押金;column:deposit;size:10;"`               //押金
	Status     *string    `json:"status" form:"status" gorm:"comment:订单状态("待处理'，活动'，'已完成'，'已取消");column:status;"` //订单状态("待处理'，活动'，'已完成'，'已取消")
	SignedAt   *time.Time `json:"signedAt" form:"signedAt" gorm:"comment:签约时间;column:signed_at;"`                 //签约时间
}

// TableName rentalOrder表 RentalOrder自定义表名 rental_order
func (RentalOrder) TableName() string {
	return "rental_order"
}
