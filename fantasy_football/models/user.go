package models

import (
	"database/sql"
	"fmt"
	"math/rand"
	"patrickjr/fantasy_football/lib/ff_random"
	"patrickjr/fantasy_football/models/model_constants"

	"patrickjr/fantasy_football/models/validate_user"
	"time"

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
		return returningUserWithCookie(email, cookie)
	} else {
		return nil
	}
}

func UserLogin(m map[string]string) *User {
	user, err := returningUser(m)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return user
}

func UserRegister(m map[string]string) (*User, error) {
	user, err := newUser(m)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func UserSetSession(u *User) {

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
	return &User{UserName: m[constants.Name], Email: m[constants.Email], Cookie: generate_cookie(), Remember: false}, nil
}

func returningUser(m map[string]string) (*User, error) {
	constants := model_constants.SetUserConstants()
	err := validate_user.Login(m[constants.Email], m[constants.Password])
	if err != nil {
		return nil, err
	}
	return login_user(m[constants.Email], m[constants.Password])
}

func login_user(email string, password string) (*User, error) {
	row, err := db.LoginUserByEmail(email, password)
	return get_user(row, err)
}

func get_user(row *sql.Row, err error) (*User, error) {
	if err != nil {
		return nil, err
	}
	user := scan_user(row)
	return user, err
}

func returningUserWithCookie(email string, cookie string) *User {
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

func findUser(email string) (*User, error) {
	row, err := db.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}
	user := scan_user(row)
	return user, err
}

func get_login_cookies(s *sessions.Session) (string, string) {
	email := convert_to_string(s.Values["ff_email"])
	cookie := convert_to_string(s.Values["ff_snicker_doodle"])
	return email, cookie
}

func convert_to_string(val interface{}) string {
	if str, ok := val.(string); ok {
		return str
	} else {
		return ""
	}
}

// stmt, err := d.Prepare(`SELECT user_name, email, cookie, ip,  FROM users WHERE email = $1`)
func scan_user(row *sql.Row) *User {
	var name string
	var email string
	err := row.Scan(&name, &email)
	if err != nil {
		return nil
	}
	return &User{name, email, generate_cookie(), false}
}

func generate_cookie() string {
	var src = rand.NewSource(time.Now().UnixNano())
	return ff_random.RandStringBytesMaskImprSrc(49, src)
}
