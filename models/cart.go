package models

import "time"

type Cart struct {
	ID       int       `json:"id"`
	BookID   int       `json:"book_id" gorm:"type: int"`
	Book     Book      `json:"book" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TransactionID int `json:"transaction_id"`
	Transaction Transaction `json:"-" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Total		 int			 `json:"total"`
	CreateAt time.Time `json:"-"`
}