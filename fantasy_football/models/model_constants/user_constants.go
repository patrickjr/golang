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

const (
	PasswordLength  = "reg_err_0"
	PasswordNoMatch = "reg_err_1"
	EmailNameTaken  = "reg_err_2"
	InvalidEmail    = "reg_err_3"

	InvalidLogin = "log_err_0"
)

/*
	registration errors
	reg_err_0 = "password not long enough"
	reg_err_1 = "password doesn't match"
	reg_err_2 = "email/username taken"

*/
