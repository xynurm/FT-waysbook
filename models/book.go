package models

import "time"

type Book struct {
	ID              int    `json:"id"`
	Title           string `json:"title" gorm:"type: varchar(50)"`
	PublicationDate time.Time	`json:"publication_date"`
	Pages 					int `json:"pages" gorm:"type:int"`
	ISBN						int `json:"isbn" gorm:"type: int"`
	Price						int `json:"price" gorm:"type: int"`
	Author					string `json:"author" gorm:"type:varchar(50)"`
	Description			string `json:"description" gorm:"type: text"`
	Image						string `json:"image" gorm:"type: varchar(255)"`
	FilePDF					string `json:"filePDF" gorm:"type: varchar(255)"`
	Status					string `json:"status" gorm:"type: varchar(50)"`
	CreateAt				time.Time	`json:"-"`
	UpdateAt				time.Time `json:"-"`
}