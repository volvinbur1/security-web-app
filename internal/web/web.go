package web

import (
	"fmt"
	"github.com/volvinbur1/security-web-app/internal/auth"
	"github.com/volvinbur1/security-web-app/internal/cmn"
	"github.com/volvinbur1/security-web-app/internal/db"
	"github.com/volvinbur1/security-web-app/internal/routing"
	"log"
	"net/http"
)

type Worker struct {
	dbManager *db.Manager
}

func NewWorker() *Worker {
	dbMgr := db.New()
	return &Worker{dbManager: dbMgr}
}

func (w *Worker) Stop() {
	w.dbManager.Disconnect()
}

func (w *Worker) RegistrationHandler(rw http.ResponseWriter, req *http.Request) {
	routing.RegisterPage(rw, req)

	if req.Method == "POST" {
		err := req.ParseForm()
		if err != nil {
			log.Fatal(err)
		}

		newUser := cmn.User{}
		newUser.Login = req.FormValue("email")
		newUser.Password = req.FormValue("psw")
		newUser.Name = req.FormValue("name")
		newUser.Surname = req.FormValue("lastname")

		err = auth.Register(w.dbManager, newUser)
		if err != nil {
			fmt.Fprintf(rw, err.Error())
		} else {
			fmt.Fprintf(rw, "User registered")
		}
	}
}

func (w *Worker) LoginHandler(rw http.ResponseWriter, req *http.Request) {
	routing.LoginPage(rw, req)

	if req.Method == "POST" {
		err := req.ParseForm()
		if err != nil {
			log.Fatal(err)
		}

		loggingUser := cmn.User{}
		loggingUser.Login = req.FormValue("email")
		loggingUser.Password = req.FormValue("psw")

		err = auth.LoginUser(w.dbManager, loggingUser)
		if err != nil {
			fmt.Fprintf(rw, err.Error())
		} else {
			fmt.Fprintf(rw, "User logged in")
		}
	}
}
