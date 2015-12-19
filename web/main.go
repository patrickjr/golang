package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

type webApplication struct {
	templates *template.Template
}

func NewWebApplication() *webApplication {
	webApp := new(webApplication)
	webApp.init()
	return webApp
}

func (webApp *webApplication) init() {
	webApp.templates = template.Must(template.ParseFiles("index.html"))
}

func (app *webApplication) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	app.renderTemplate(w, "index")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func (app *webApplication) renderTemplate(w http.ResponseWriter, tmpl string) {
	err := app.templates.ExecuteTemplate(w, tmpl+".html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(pwd)
	webApp := NewWebApplication()
	router := httprouter.New()

	router.ServeFiles("/assets/*filepath", http.Dir("assets"))

	router.GET("/", webApp.Index)
	// router.GET("/hello/:name", Hello)
	log.Fatal(http.ListenAndServe(":8080", router))
}
