package models

import (
	"net/http"
	"patrickjr/fantasy_football/models/model_constants"
	"patrickjr/fantasy_football/models/validate_user"
)

type User struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Remember bool   `json:"remember"`
}

func UserLogin(r *http.Request) *User {
	return &User{}
}

func UserRegister(m map[string]string) *User {
	user, err := newUser(m)
	if err != nil {
		return nil
	}
	return user
}

// creates a new user in the database, after validation
func newUser(m map[string]string) (*User, error) {
	constants := model_constants.SetUserConstants()
	err := validate_user.Email(m[constants.Email])
	err = validate_user.Password(m[constants.Password], m[constants.Confirm_Password])
	if err != nil {
		return nil, err
	} else {
		return createUser(m)
	}
}

func createUser(m map[string]string) (*User, error) {
	_, err := db.CreateUser(m)
	if err != nil {
		return nil, err
	}
	constants := model_constants.SetUserConstants()
	return &User{UserName: m[constants.Name], Email: m[constants.Email], Remember: false}, nil
}
