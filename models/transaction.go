package models

import "time"

type Transaction struct {
	ID       int    `json:"id"`
	UserID   int    `json:"user_id" gorm:"type: int"`
	User     User   `json:"user"`
	Cart     []Cart `json:"cart"`
	Total		 int		`json:"total" gorm:"type: int"`
	Status	 string	`json:"status" gorm:"type: varchar(20)"`
	CreateAt time.Time	`json:"-"`
	UpdateAt time.Time	`json:"-"`
}