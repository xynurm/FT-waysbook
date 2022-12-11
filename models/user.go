package models

import "time"

type User struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname" gorm:"type: varchar(50)"`
	Email 	 string `json:"email" gorm:"type: varchar(50)"`
	Password string `json:"password" gorm:"type: varchar(100)"`
	Image		 string `json:"image" gorm:"type: varchar(255)"`
	Role		 string `json:"role" gorm:"type: varchar(20)"`
	Gender   string `json:"gender" gorm:"type: varchar(20)"`
	Phone    string `json:"phone" gorm:"type :varchar(20)"`
	Address  string `json:"address" gorm:"type:text"`
	CreateAt  time.Time	`json:"-"`
	UpdateAt  time.Time	`json:"-"`
}