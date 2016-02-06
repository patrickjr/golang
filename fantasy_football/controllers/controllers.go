package controllers

import (
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, templates *template.Template, templateName string, data interface{}) {
	err := templates.ExecuteTemplate(w, templateName+".html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
