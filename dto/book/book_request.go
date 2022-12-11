package bookdto

type CreateBook struct {
	Title           string `json:"title" form:"title" validate:"required"`
	PublicationDate string `json:"publication_date" form:"publicationDate" validate:"required"`
	Pages           int    `json:"pages" form:"pages" validate:"required"`
	ISBN            int    `json:"isbn" form:"isbn" validate:"required"`
	Price           int    `json:"price" form:"price" validate:"required"`
	Author          string `json:"author" form:"author" validate:"required"`
	Description     string `json:"description" form:"description" validate:"required"`
	Image           string `json:"image" form:"image"`
	FilePDF         string `json:"filePDF" form:"filePDF"`
	Status          string `json:"status" `
}

type UpdateBook struct {
	Title           string `json:"title" form:"title"`
	PublicationDate string `json:"publication_date" form:"publicationDate"`
	Pages           int    `json:"pages" form:"pages"`
	ISBN            int    `json:"isbn" form:"isbn"`
	Price           int    `json:"price" form:"price"`
	Author          string `json:"author" form:"author"`
	Description     string `json:"description" form:"description"`
	Image           string `json:"image" form:"image"`
	FilePDF         string `json:"filePDF" form:"filePDF"`
	Status          string `json:"status" form:"status"`
}

type UpdateBookPromo struct {
	Status string `json:"status"`
}