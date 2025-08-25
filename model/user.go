package model

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var ValidUser = User{
	Username: "test",
	Password: "testpass",
}
