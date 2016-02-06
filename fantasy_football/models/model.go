package models

import (
	"database/sql"
	"fmt"
	"log"
	"patrickjr/fantasy_football/environment"
	"patrickjr/fantasy_football/models/model_constants"

	_ "github.com/lib/pq"
)

type DbWrapper struct {
	*sql.DB
}

var db *DbWrapper

func Init() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", environment.DB_USER, environment.DB_PASSWORD, environment.DB_NAME)
	database, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
	}
	db = &DbWrapper{database}
}

func (d *DbWrapper) CreateUser(m map[string]string) (sql.Result, error) {
	constants := model_constants.SetUserConstants()
	stmt, err := d.Prepare(`INSERT INTO users (user_name, email, password, forgot_password)
                           VALUES            ($1       ,    $2,       $3,             $4)`)
	checkErr(err)
	res, err := stmt.Exec(m[constants.Name], m[constants.Email], m[constants.Password], false)
	return res, err
}

func (d *DbWrapper) FindUserByName(name string) (*sql.Row, error) {
	stmt, err := d.Prepare(`SELECT user_name, email FROM users WHERE user_name = $1`)
	checkErr(err)

	row := stmt.QueryRow(name)
	return row, err
}

func (d *DbWrapper) FindUserByEmail(email string) (*sql.Row, error) {
	stmt, err := d.Prepare(`SELECT user_name, email FROM users WHERE email = $1 `)
	checkErr(err)
	row := stmt.QueryRow(email)
	return row, err
}

func (d *DbWrapper) LoginUserByEmail(email string, password string) (*sql.Row, error) {
	stmt, err := d.Prepare(`SELECT user_name, email FROM users WHERE email = $1 AND password = $2 `)
	checkErr(err)
	row := stmt.QueryRow(email, password)
	return row, err
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

/* used for initializing database for testing purposes only*/
func InitTest() *DbWrapper {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", environment.DB_USER, environment.DB_PASSWORD, environment.DB_NAME)
	database, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
	}
	return &DbWrapper{database}
}
