package usersdto

type UpdateUser struct {
	Fullname string `json:"fullname" form:"fullname"`
	Password string `json:"password" form:"password"`
	Image    string `json:"image" form:"image"`
	Gender   string `json:"gender" form:"gender"`
	Phone    string `json:"phone" form:"phone"`
	Address  string `json:"address" form:"address"`
}