package ff_crypto

import (
	"errors"
	"patrickjr/fantasy_football/lib/ff_utility"

	"golang.org/x/crypto/bcrypt"
)

func DigestPassword(password string) (string, error) {
	ps := []byte(password)
	digest, err := bcrypt.GenerateFromPassword(ps, 10)
	if err != nil {
		return "error", err
	}
	password = ff_utility.BytesToString(digest)
	return password, nil
}

func ComparePassword(password string, hash string) error {
	ps := []byte(password)
	hs := []byte(hash)
	err := bcrypt.CompareHashAndPassword(hs, ps)
	if err != nil {
		return errors.New("invalid combination")
	}
	return nil
}
