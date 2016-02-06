package validate_user

import (
	"errors"
	"patrickjr/fantasy_football/models/model_constants"
	"regexp"
)

func Email(email string) error {
	if len(email) == 0 {
		return errors.New(model_constants.InvalidEmail)
	}
	emailRegex := `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
	re := regexp.MustCompile(emailRegex)
	value := re.MatchString(email)
	if !value {
		return errors.New(model_constants.InvalidEmail)
	}
	return nil
}

func Password(pass1 string, pass2 string) error {
	if (pass1 == pass2) && (len(pass1) > 6) {
		return nil
	} else {
		if len(pass1) <= 6 {
			return errors.New(model_constants.PasswordLength)
		}
		return errors.New(model_constants.PasswordNoMatch)
	}
}

func Login(email string, pass string) error {
	if (len(email) > 0) && (len(pass) > 0) {
		if (len(email) < 40) && (len(pass) < 40) {
			return nil
		}
	}
	return errors.New(model_constants.InvalidLogin)
}
