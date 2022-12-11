package transactiondto

type Transaction struct {
	Total int `json:"total" validate:"required"`
}
