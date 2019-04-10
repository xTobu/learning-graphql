package models

import (
	"github.com/jinzhu/gorm"
)

// User struct
type User struct {
	ID       int    `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
}

// GetUsers 取得資料表：users 的所有資料
func GetUsers() ([]*User, error) {
	var users []*User
	err := db.Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return users, nil
}
