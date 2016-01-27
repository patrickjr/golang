package controllers

import (
	"encoding/json"
	"html/template"
	"net/http"
	"patrickjr/fantasy_football/models"
	"patrickjr/fantasy_football/models/model_constants"

	"github.com/julienschmidt/httprouter"
)

func Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params, templates *template.Template) {
	if r.Method == "POST" {
		details := parseRegisterForm(r)
		data := models.UserRegister(details)
		json, err := json.Marshal(data)
		if err != nil {

		}
		w.Write(json)
	}
}

func parseRegisterForm(r *http.Request) map[string]string {
	r.ParseForm()
	constants := model_constants.SetUserConstants()
	details := map[string]string{
		constants.Name:             r.FormValue(constants.Name),
		constants.Email:            r.FormValue(constants.Email),
		constants.Password:         r.FormValue(constants.Password),
		constants.Confirm_Password: r.FormValue(constants.Confirm_Password),
	}
	return details
}
