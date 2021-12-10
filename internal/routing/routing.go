package routing

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

type RegisterStruct struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

func LoginPage(rw http.ResponseWriter, req *http.Request) {
	path := filepath.Join("web", "app", "html", "loginPage.html")

	tmpl, err := template.ParseFiles(path)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	err = tmpl.Execute(rw, nil)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

}
func RegisterPage(rw http.ResponseWriter, req *http.Request) {

	path := filepath.Join("web", "app", "html", "registerPage.html")

	tmpl, err := template.ParseFiles(path)

	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
	err = tmpl.Execute(rw, nil)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
	if req.Method == "POST" {
		var registerStruct RegisterStruct

		err = req.ParseForm()
		if err != nil {
			log.Fatal(err)
		}

		registerStruct.Email = req.FormValue("email")
		password := req.FormValue("psw")
		/*
			some password hashing

		*/
		registerStruct.Password = password
		fmt.Println(registerStruct)
	} else {
		//some error
	}

}
