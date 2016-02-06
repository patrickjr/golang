package controllers

import (
	_ "encoding/json"
	"net/http"
	"patrickjr/fantasy_football/models"
	"patrickjr/fantasy_football/models/model_constants"

	"patrickjr/fantasy_football/web_app/web_app_interface"

	"github.com/julienschmidt/httprouter"
)

func Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params, app web_app_interface.Interface) {
	if r.Method == "POST" {
		details := parseRegisterForm(r)
		data, err := models.UserRegister(details)
		if err != nil {
			app.FlashMessages(w, r, err.Error())
			app.RouteSignUp(w, r)
		} else {
			app.CreateNewUserSession(w, r, data)
			app.RouteHome(w, r)
		}
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
