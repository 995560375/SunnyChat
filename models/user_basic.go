package models

import "github.com/jinzhu/gorm"

type UserBasic struct {
	gorm.Model
	Name          string
	PassWord      string
	Phone         string
	Email         string
	Identity      string
	ClientIp      string
	ClientPort    string
	LoginTime     uint64
	HeartBeatTime uint64
	LogOutTime    uint64
	IsLogOut      bool
	DeviceInfo    string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}
