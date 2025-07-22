package house

import (
	"time"
)

type House struct {
	HouseID                 int64      `json:"houseId" gorm:"primaryKey;column:house_id;autoIncrement;comment:房源ID"`
	Title                   string     `json:"title" gorm:"type:varchar(100);not null;comment:房源标题"`
	Description             string     `json:"description" gorm:"type:text;comment:房源描述"`
	LandlordID              int64      `json:"landlordId" gorm:"not null;comment:发布人ID"`
	Address                 string     `json:"address" gorm:"type:varchar(255);not null;comment:详细地址"`
	RegionID                *int64     `json:"regionId" gorm:"comment:区域/小区ID"`
	CommunityID             *int64     `json:"communityId" gorm:"comment:小区ID"`
	Price                   float64    `json:"price" gorm:"type:decimal(10,2);not null;comment:价格"`
	Area                    *float64   `json:"area" gorm:"comment:面积"`
	Layout                  *string    `json:"layout" gorm:"type:varchar(50);comment:户型"`
	Floor                   *string    `json:"floor" gorm:"type:varchar(20);comment:楼层"`
	OwnershipCertificateURL string     `json:"ownershipCertificateUrl" gorm:"type:varchar(255);not null;comment:产权证明图片"`
	Orientation             *string    `json:"orientation" gorm:"type:varchar(20);comment:朝向"`
	Decoration              *string    `json:"decoration" gorm:"type:varchar(50);comment:装修"`
	Facilities              *string    `json:"facilities" gorm:"type:varchar(255);comment:配套设施（逗号分隔）"`
	Status                  string     `json:"status" gorm:"type:enum('active','inactive','rented');default:'active';not null;comment:状态（活跃，不活跃，已租用）"`
	CreatedAt               time.Time  `json:"createdAt" gorm:"autoCreateTime;comment:发布时间"`
	UpdatedAt               time.Time  `json:"updatedAt" gorm:"autoUpdateTime;comment:更新时间"`
	DeletedAt               *time.Time `json:"deletedAt" gorm:"index;comment:删除时间"`
}

func (h *House) TableName() string {
	return "house"
}
