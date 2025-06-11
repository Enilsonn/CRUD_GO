package handlers

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type BD struct {
	Users map[int]User
	Len   int
}
