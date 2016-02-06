package controllers

import (
	"net/http"
	"patrickjr/fantasy_football/web_app/web_app_interface"

	"github.com/julienschmidt/httprouter"
)

func Logout(w http.ResponseWriter, r *http.Request, _ httprouter.Params, app web_app_interface.Interface) {
	session := app.GetSession(w, r)
	app.RemoveSession(w, r, session)
	app.RouteIndex(w, r)
}
