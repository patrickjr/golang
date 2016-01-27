package controllers

import (
	_ "encoding/json"
	"html/template"
	"net/http"
	"patrickjr/fantasy_football/models"

	"github.com/julienschmidt/httprouter"
)

type loginDetails struct {
	Email string `json:"email"`
	Phone string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params, templates *template.Template) {
	models.UserLogin(r)
}
