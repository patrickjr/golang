package main

import (
	"log"
	"net/http"
	"patrickjr/fantasy_football/models"
	"patrickjr/fantasy_football/web_app"

	"github.com/julienschmidt/httprouter"
)

func main() {
	models.Init()
	webApp := web_app.NewWebApplication()
	router := httprouter.New()
	router.ServeFiles("/assets/*filepath", http.Dir("assets"))
	router.GET("/", webApp.Index)
	router.POST("/login", webApp.Login)
	router.POST("/register", webApp.Register)
	router.GET("/sign_in", webApp.SignIn)
	router.GET("/sign_up", webApp.SignUp)
	router.GET("/logout", webApp.Logout)
	router.GET("/home", webApp.Home)
	log.Fatal(http.ListenAndServe(":8080", router))
}
