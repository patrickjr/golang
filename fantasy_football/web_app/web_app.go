/* package web_app
 * this package is responsible for routing requests to respective controllers
 * it is also responsible for initializing and passing templates to the controllers
 * so that the controllers can display/write the correct template to the user/client
 */

package web_app

import (
	"encoding/gob"
	"html/template"
	"net/http"
	"patrickjr/fantasy_football/models"

	"github.com/gorilla/sessions"
)

type M map[string]interface{}

type webApplication struct {
	templates *template.Template
	store     *sessions.CookieStore
}

func (webApp *webApplication) ServeHTTP(w http.ResponseWriter, req *http.Request) {}

func NewWebApplication() *webApplication {
	webApp := new(webApplication)
	webApp.init()
	return webApp
}

func (webApp *webApplication) init() {
	webApp.store = sessions.NewCookieStore([]byte("something-very-secret"))
	webApp.templates = template.Must(template.ParseGlob("views/*")).Funcs(
		template.FuncMap{
			"eq": func(x, b bool) bool {
				return x == b
			},
		})

	gob.Register(&models.User{})
	gob.Register(&M{})
}

func (webApp *webApplication) get_session(r *http.Request) (*sessions.Session, error) {
	return webApp.store.Get(r, "something-very-secret")
}

func (webApp *webApplication) GetTemplates() *template.Template {
	return webApp.templates
}

func (webApp *webApplication) GetStore() *sessions.CookieStore {
	return webApp.store
}

func (webApp *webApplication) IsNewSession(w http.ResponseWriter, r *http.Request) bool {
	session, err := webApp.get_session(r)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return true
	}
	return session.IsNew
}

func (webApp *webApplication) GetSession(w http.ResponseWriter, r *http.Request) *sessions.Session {
	session, err := webApp.get_session(r)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return nil
	}
	return session
}

func (webApp *webApplication) RemoveSession(w http.ResponseWriter, r *http.Request, s *sessions.Session) {
	s.Options.MaxAge = -1
	s.Save(r, w)
}

func (webApp *webApplication) SaveSession(w http.ResponseWriter, r *http.Request, s *sessions.Session) {
	s.Save(r, w)
}

func (webApp *webApplication) CreateNewUserSession(w http.ResponseWriter, r *http.Request, user *models.User) {
	session := webApp.GetSession(w, r)
	if session != nil {
		session.Values["ff_user"] = user
		session.Save(r, w)
	}
}

func (webApp *webApplication) FlashMessages(w http.ResponseWriter, r *http.Request, msg string) {
	s := webApp.GetSession(w, r)
	s.AddFlash(msg)
	s.Save(r, w)
}

func (webApp *webApplication) Data(w http.ResponseWriter, r *http.Request, messages map[string]interface{}) interface{} {
	var m = make(map[string]interface{})
	if len(messages) > 0 {
		m = messages
	}
	s := webApp.GetSession(w, r)
	if flashes := s.Flashes(); len(flashes) > 0 {
		m["Flashes"] = flashes
	}
	s.Save(r, w)
	return m
}

func (webApp *webApplication) UserInfo(w http.ResponseWriter, r *http.Request) *models.User {
	session := webApp.GetSession(w, r)
	val := session.Values["ff_user"]
	if _, ok := val.(*models.User); !ok {
		// Handle the case that it's not an expected type
		// throw a server 500 error
		return nil
	} else {
		return val.(*models.User)
	}
}

/* routes for interacting with the web_app interface
 * encapsulates the directories and redirect codes from the caller
 */

func (webApp *webApplication) RouteHome(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/home", 302)
}

func (webApp *webApplication) RouteIndex(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", 302)
}

func (webApp *webApplication) RouteSignIn(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/sign_in", 302)
}

func (webApp *webApplication) RouteSignUp(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/sign_up", 302)
}
