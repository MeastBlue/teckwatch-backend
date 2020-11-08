package model

type User struct {
	ID        string `json:"id"`
	Lastname  string `json:"lastname"`
	Firstname string `json:"firstname"`
	Surname   string `json:"surname"`
	Avatar    string `json:"avatar"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type NewUser struct {
	Lastname  string `json:"lastname"`
	Firstname string `json:"firstname"`
	Surname   string `json:"surname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
