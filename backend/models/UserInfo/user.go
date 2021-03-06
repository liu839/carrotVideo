package UserInfo

import "github.com/jinzhu/gorm"

//用户信息表
type User struct {
	gorm.Model
	UserName  string `gorm:"unique;not null;index:addr"`
	Password  string `gorm:"not null"`
	Name      string
	Age       *int32
	Gender    int32 `gorm:"default:0"`
	Lever     *int32
	Address   string
	Phone     string `gorm:"unique"`
	StarCount int32  //点赞数
}
