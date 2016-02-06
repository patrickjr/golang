package web_app_interface

import (
	"html/template"
	"net/http"
	"patrickjr/fantasy_football/models"

	"github.com/gorilla/sessions"
)

type Interface interface {
	GetTemplates() *template.Template
	GetStore() *sessions.CookieStore
	IsNewSession(http.ResponseWriter, *http.Request) bool
	GetSession(http.ResponseWriter, *http.Request) *sessions.Session
	RemoveSession(http.ResponseWriter, *http.Request, *sessions.Session)
	SaveSession(http.ResponseWriter, *http.Request, *sessions.Session)
	CreateNewUserSession(http.ResponseWriter, *http.Request, *models.User)

	UserInfo(http.ResponseWriter, *http.Request) *models.User

	/* flash messages provided to templates */
	FlashMessages(http.ResponseWriter, *http.Request, string)
	Data(http.ResponseWriter, *http.Request, map[string]interface{}) interface{}

	/* routes */
	RouteHome(http.ResponseWriter, *http.Request)
	RouteIndex(http.ResponseWriter, *http.Request)
	RouteSignIn(http.ResponseWriter, *http.Request)
	RouteSignUp(http.ResponseWriter, *http.Request)
}
