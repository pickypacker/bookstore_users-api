package users

type User struct {
	Id          int64  `json:"id"`
	FristName   string `1json:"first_name"`
	LastName    string `json:"lst_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}
