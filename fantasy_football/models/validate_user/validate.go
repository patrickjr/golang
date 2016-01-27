package validate_user

import (
	"errors"
	"regexp"
)

func Email(email string) error {
	if len(email) == 0 {
		return errors.New("email cannot be empty")
	}
	emailRegex := `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
	re := regexp.MustCompile(emailRegex)
	value := re.MatchString(email)
	if !value {
		return errors.New("email is invalid")
	}
	return nil
}

func Password(pass1 string, pass2 string) error {
	if (pass1 == pass2) && (len(pass1) > 6) {
		return nil
	} else {
		return errors.New("password is invalid")
	}
}
