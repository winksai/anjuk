package order

import "time"

// Contract 表示电子合同表结构
type Contract struct {
	ContractID  int64     `json:"contract_id" gorm:"primaryKey;autoIncrement;comment:合同ID"`
	OrderID     int64     `json:"order_id" gorm:"not null;comment:关联订单ID"`
	ContractURL string    `json:"contract_url" gorm:"type:varchar(255);not null;comment:合同文件URL"`
	SignTime    time.Time `json:"sign_time" gorm:"type:datetime;comment:签约时间"`
	Status      string    `json:"status" gorm:"type:enum('signed','unsigned','expired');not null;default:'unsigned';comment:合同状态（已签署、未签署、已过期）"`
}

// TableName 指定Contract结构体对应的数据库表名
func (c *Contract) TableName() string {
	return "contract"
}
