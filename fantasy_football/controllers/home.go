package controllers

import (
	_ "encoding/json"
	"net/http"
	"patrickjr/fantasy_football/web_app/web_app_interface"

	"github.com/julienschmidt/httprouter"
)

func Home(w http.ResponseWriter, r *http.Request, _ httprouter.Params, app web_app_interface.Interface) {
	user := app.UserInfo(w, r)
	if user == nil {
		app.RouteIndex(w, r)
	} else {
		data := make(map[string]interface{})
		data["user"] = app.UserInfo(w, r)
		RenderTemplate(w, app.GetTemplates(), "home", app.Data(w, r, data))
	}
}
