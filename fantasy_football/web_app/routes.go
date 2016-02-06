package web_app

import (
	"net/http"
	"patrickjr/fantasy_football/controllers"

	"github.com/julienschmidt/httprouter"
)

func (app *webApplication) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	controllers.Index(w, r, nil, app)
}

func (app *webApplication) Register(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	controllers.Register(w, r, ps, app)
}

func (app *webApplication) Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	controllers.Login(w, r, ps, app)
}

func (app *webApplication) SignUp(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	controllers.SignUp(w, r, ps, app)
}

func (app *webApplication) SignIn(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	controllers.SignIn(w, r, ps, app)
}

func (app *webApplication) Home(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	controllers.Home(w, r, ps, app)
}

func (app *webApplication) Logout(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	controllers.Logout(w, r, ps, app)
}
