package user

import (
	"time"
)

type UserBase struct {
	UserID     int64      `json:"userId" gorm:"primaryKey;column:user_id;autoIncrement;comment:用户唯一ID"`
	Name       string     `json:"name" gorm:"type:varchar(50);not null;comment:用户昵称/姓名"`
	RealName   *string    `json:"realName" gorm:"type:varchar(30);comment:真实姓名"`
	Phone      string     `json:"phone" gorm:"type:char(11);not null;comment:手机号"`
	Password   string     `json:"password" gorm:"type:char(32);not null;comment:密码（加密存储）"`
	Avatar     *string    `json:"avatar" gorm:"type:text;comment:头像URL"`
	RoleID     int64      `json:"roleId" gorm:"not null;comment:角色id"`
	Sex        *string    `json:"sex" gorm:"type:enum('男','女');comment:用户性别"`
	RealStatus *int8      `json:"realStatus" gorm:"comment:用户实名状态(1:已实名2:未实名)"`
	Status     int8       `json:"status" gorm:"default:1;not null;comment:状态（0禁用1正常）"`
	CreatedAt  time.Time  `json:"createdAt" gorm:"autoCreateTime;comment:注册时间"`
	UpdatedAt  time.Time  `json:"updatedAt" gorm:"autoUpdateTime;comment:更新时间"`
	DeletedAt  *time.Time `json:"deletedAt" gorm:"index;comment:删除时间"`
}

func (UserBase) TableName() string {
	return "user_base"
}
