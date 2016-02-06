package controllers

import (
	_ "encoding/json"
	"net/http"
	"patrickjr/fantasy_football/web_app/web_app_interface"

	"github.com/julienschmidt/httprouter"
)

func SignIn(w http.ResponseWriter, r *http.Request, _ httprouter.Params, app web_app_interface.Interface) {
	user := app.UserInfo(w, r)
	data := make(map[string]interface{})
	if user == nil {
		RenderTemplate(w, app.GetTemplates(), "sign_in", app.Data(w, r, data))
	} else {
		app.RouteHome(w, r)
	}
}
