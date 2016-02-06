package models

import (
	"database/sql"
	"errors"
	"fmt"
	"patrickjr/fantasy_football/lib/ff_crypto"
	"patrickjr/fantasy_football/lib/ff_utility"
	"patrickjr/fantasy_football/models/model_constants"
	"patrickjr/fantasy_football/models/validate_user"

	"github.com/gorilla/sessions"
)

type User struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Cookie   string `json:"cookie"`
	Remember bool   `json:"remember"`
}

func UserLoginWithSession(s *sessions.Session) *User {
	email, cookie := get_login_cookies(s)
	if cookie != "" && email != "" {
		return returning_user_with_cookie(email, cookie)
	} else {
		return nil
	}
}

func UserLogin(m map[string]string) (*User, error) {
	user, err := returning_user(m)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func UserRegister(m map[string]string) (*User, error) {
	user, err := new_user(m)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func UserSetSession(u *User) {

}

// creates a new user in the database, after validation
func new_user(m map[string]string) (*User, error) {
	constants := model_constants.SetUserConstants()
	err := validate_user.Email(m[constants.Email])
	err = validate_user.Password(m[constants.Password], m[constants.Confirm_Password])
	if err != nil {
		return nil, err
	}
	m, err = digest_password(m)
	if err != nil {
		return nil, err
	}
	return create_user(m)
}

func create_user(m map[string]string) (*User, error) {
	_, err := db.CreateUser(m)
	if err != nil {
		return nil, errors.New("reg_err_2")
	}
	constants := model_constants.SetUserConstants()
	return &User{UserName: m[constants.Name], Email: m[constants.Email], Cookie: ff_utility.GenerateCookie(), Remember: false}, nil
}

func returning_user(m map[string]string) (*User, error) {
	constants := model_constants.SetUserConstants()
	err := validate_user.Login(m[constants.Email], m[constants.Password])
	if err != nil {
		return nil, err
	}
	return login_user(m[constants.Email], m[constants.Password])
}

func login_user(email string, password string) (*User, error) {
	row, err := db.LoginUserByEmail(email)
	if valid_password(password, row, err) {
		row, err = db.FindUserByEmail(email)
		return get_user(row, err)
	} else {
		return nil, errors.New(model_constants.InvalidLogin)
	}
}

func valid_password(password string, row *sql.Row, err error) bool {
	if err != nil {
		return false
	}
	var hashedPassword string
	row.Scan(&hashedPassword)
	err = ff_crypto.ComparePassword(password, hashedPassword)
	if err != nil {
		return false
	} else {
		return true
	}
}

func get_user(row *sql.Row, err error) (*User, error) {
	if err != nil {
		return nil, errors.New(model_constants.InvalidLogin)
	}
	user := scan_user(row)
	return user, err
}

func returning_user_with_cookie(email string, cookie string) *User {
	data, err := db.FindUserByEmail(email)
	if err != nil {
		return nil
	}
	user := scan_user(data)
	if user.Cookie == cookie {
		return user
	}
	return nil
}

func find_user(email string) (*User, error) {
	row, err := db.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}
	user := scan_user(row)
	return user, err
}

// stmt, err := d.Prepare(`SELECT user_name, email, cookie, ip,  FROM users WHERE email = $1`)
func scan_user(row *sql.Row) *User {
	var name string
	var email string
	err := row.Scan(&name, &email)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &User{name, email, ff_utility.GenerateCookie(), false}
}

func digest_password(m map[string]string) (map[string]string, error) {
	constants := model_constants.SetUserConstants()
	var err error
	m[constants.Password], err = ff_crypto.DigestPassword(m[constants.Password])
	return m, err
}

func get_login_cookies(s *sessions.Session) (string, string) {
	email := ff_utility.ConvertToString(s.Values["ff_email"])
	cookie := ff_utility.ConvertToString(s.Values["ff_snicker_doodle"])
	return email, cookie
}
