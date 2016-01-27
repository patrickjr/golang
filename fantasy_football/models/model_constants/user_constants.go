package model_constants

type User struct {
	Name             string
	Email            string
	Password         string
	Confirm_Password string
}

func SetUserConstants() *User {
	return &User{
		Name:             "user[name]",
		Email:            "user[email]",
		Password:         "user[password]",
		Confirm_Password: "user[confirm_password]",
	}
}
