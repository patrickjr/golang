package models

import (
	"patrickjr/fantasy_football/models"
	"testing"
)

func TestCreateUser(t *testing.T) {
	db := models.InitTest()
	details := map[string]string{
		"user_name":        "patrick2",
		"email":            "tupac2@gmail.com",
		"password":         "password",
		"confirm_password": "password",
	}
	db.CreateUser(details)
	t.Log(db.CreateUser(details))
}

func TestFindUserByEmail(t *testing.T) {
	db := models.InitTest()
	row, err := db.FindUserByEmail("tupac2@gmail.com")
	if err != nil {
		t.Error(err)
	}
	var name string
	var email string
	row.Scan(&name, &email)
	if email != "tupac2@gmail.com" || name != "patrick2" {
		t.Error(name + " " + email)
	}
}

func TestFindUserByName(t *testing.T) {
	db := models.InitTest()
	row, err := db.FindUserByName("patrick2")
	if err != nil {
		t.Error(err)
	}
	var name string
	var email string
	row.Scan(&name, &email)
	if email != "tupac2@gmail.com" || name != "patrick2" {
		t.Error(name + " " + email)
	}
}
