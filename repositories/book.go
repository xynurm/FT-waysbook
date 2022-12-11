package repositories

import (
	"waysbook/models"

	"gorm.io/gorm"
)

type BookRepository interface {
	CreateBook(book models.Book) (models.Book, error)
	FindBook() ([]models.Book, error)
	GetBook(ID int) (models.Book, error)
	UpdateBook(book models.Book) (models.Book, error)
	DeleteBook(book models.Book) (models.Book, error)
	FindBookPromo() ([]models.Book, error)
	FindBookRegular() ([]models.Book, error)
	UpdateBookPromo(book models.Book) (models.Book, error)
}

func RepositoryBook(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateBook(book models.Book) (models.Book, error) {
	err := r.db.Create(&book).Error
	return book, err
}

func (r *repository) FindBook() ([]models.Book, error) {
	var books []models.Book
	err := r.db.Find(&books).Error
	return books, err
}

func (r *repository) GetBook(ID int) (models.Book, error) {
	var book models.Book
	err := r.db.First(&book, ID).Error
	return book, err
}

func (r *repository) UpdateBook(book models.Book) (models.Book, error) {
	err := r.db.Save(&book).Error
	return book, err
}

func (r *repository) DeleteBook(book models.Book) (models.Book, error) {
	err := r.db.Delete(&book).Error
	return book, err
}

func (r *repository) FindBookPromo() ([]models.Book, error) {
	var books []models.Book
	err := r.db.Where("status= ?", "promo").Find(&books).Error
	return books, err
}

func (r *repository) FindBookRegular() ([]models.Book, error) {
	var books []models.Book
	err := r.db.Where("status= ?", "regular").Find(&books).Error
	return books, err
}

func (r *repository) UpdateBookPromo(book models.Book) (models.Book, error) {
	err := r.db.Save(&book).Error
	return book, err
}