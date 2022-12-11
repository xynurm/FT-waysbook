package bookdto

type BookResponse struct {
	Title           string `json:"title" form:"title"`
	PublicationDate string `json:"publicationDate" form:"publicationDate"`
	Pages           int    `json:"pages" form:"pages"`
	ISBN            int    `json:"isbn" form:"isbn"`
	Price           int    `json:"price" form:"price"`
	Description     string `json:"description" form:"description"`
	Image           string `json:"image" form:"image"`
	FilePDF         string `json:"filePDF" form:"filePDF"`
}