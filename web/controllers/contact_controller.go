package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"net/smtp"
	"patrickjr/web/environment"

	"github.com/julienschmidt/httprouter"
)

type contactResponse struct {
	Success bool `json:"success"`
}

type contactDetails struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Msg   string `json:"msg"`
}

func Contact(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	processForm(w, r)
}

func processForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		details := parseContactForm(r)
		if details != nil {
			sendEmail(formatEmail(details))
			response := &contactResponse{true}
			json.NewEncoder(w).Encode(response)
		} else {
			response := &contactResponse{false}
			json.NewEncoder(w).Encode(response)
		}
	}
}

func parseContactForm(r *http.Request) *contactDetails {
	name := r.FormValue("name")
	email := r.FormValue("email")
	phone := r.FormValue("phone")
	msg := r.FormValue("message")
	if name != "" && email != "" && msg != "" {
		info := &contactDetails{name, email, phone, msg}
		return info
	}
	return nil
}

func sendEmail(msg string) {
	auth := smtp.PlainAuth(
		"",
		environment.EMAIL,
		environment.EMAILPASSWORD,
		environment.SMTPSERVER,
	)
	err := smtp.SendMail(
		environment.SMTPSERVERWPORT,
		auth,
		environment.EMAIL,
		[]string{environment.EMAILRECIPIENT},
		[]byte(msg),
	)
	if err != nil {
		log.Fatal(err)
	}
}

func formatEmail(details *contactDetails) string {
	name := details.Name + "\n"
	email := details.Email + "\n"
	phone := details.Phone + "\n"
	msg := details.Msg + "\n"
	return environment.RFC822 + name + email + phone + msg + environment.CRLF
}
