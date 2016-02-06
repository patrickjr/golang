package controllers

import (
	_ "encoding/json"
	"net/http"
	"patrickjr/fantasy_football/models"
	"patrickjr/fantasy_football/models/model_constants"
	"patrickjr/fantasy_football/web_app/web_app_interface"

	"github.com/julienschmidt/httprouter"
)

func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params, app web_app_interface.Interface) {
	if r.Method == "POST" {
		details := parseLoginForm(r)
		user, err := models.UserLogin(details)
		if user != nil {
			app.CreateNewUserSession(w, r, user)
			app.RouteHome(w, r)
		} else {
			app.FlashMessages(w, r, err.Error())
			app.RouteSignIn(w, r)
		}
	}
}

func parseLoginForm(r *http.Request) map[string]string {
	r.ParseForm()
	constants := model_constants.SetUserConstants()
	details := map[string]string{
		constants.Email:    r.FormValue(constants.Email),
		constants.Password: r.FormValue(constants.Password),
	}
	return details
}
