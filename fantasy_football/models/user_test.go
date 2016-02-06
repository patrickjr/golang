package models

import (
	"patrickjr/fantasy_football/models"
	"testing"

	"patrickjr/fantasy_football/models/model_constants"
)

/* this test currently fails
 * must add digest_password before calling
 * db.CreateUser(details)
 */
func TestCreateUser(t *testing.T) {
	c := model_constants.SetUserConstants()
	db := models.InitTest()
	details := map[string]string{
		c.Name:     "patrick2",
		c.Email:    "tupac2@gmail.com",
		c.Password: "password",
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
