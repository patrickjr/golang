package web_app

import (
	"html/template"
	"net/http"
	"patrickjr/fantasy_football/controllers"

	"github.com/julienschmidt/httprouter"
)

type webApplication struct {
	templates *template.Template
}

func (webApp *webApplication) ServeHTTP(w http.ResponseWriter, req *http.Request) {}

func NewWebApplication() *webApplication {
	webApp := new(webApplication)
	webApp.init()
	return webApp
}

func (webApp *webApplication) init() {

	webApp.templates = template.Must(template.ParseGlob("views/*")).Funcs(
		template.FuncMap{
			"eq": func(x, b bool) bool {
				return x == b
			},
		})

}

func (app *webApplication) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	app.renderTemplate(w, "index")
}

func (app *webApplication) Register(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	controllers.Register(w, r, ps, app.templates)
}

func (app *webApplication) Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	controllers.Login(w, r, ps, app.templates)
}

func (app *webApplication) renderTemplate(w http.ResponseWriter, tmpl string) {
	err := app.templates.ExecuteTemplate(w, tmpl+".html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
