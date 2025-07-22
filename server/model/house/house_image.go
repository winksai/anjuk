package house

type HouseImage struct {
	ID        int64  `json:"id" gorm:"primaryKey;autoIncrement;comment:图片ID"`
	HouseID   int64  `json:"houseId" gorm:"not null;comment:房源ID"`
	ImageURL  string `json:"imageUrl" gorm:"type:varchar(255);not null;comment:图片URL"`
	SortOrder int    `json:"sortOrder" gorm:"default:0;comment:排序"`
}

// TableName 指定HouseImage结构体对应的数据库表名
func (HouseImage) TableName() string {
	return "house_image"
}
