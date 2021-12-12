package web

import (
	"github.com/volvinbur1/security-web-app/internal/cmn"
	"github.com/volvinbur1/security-web-app/internal/db"
	"github.com/volvinbur1/security-web-app/internal/web/data"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

type Worker struct {
	dbManager *db.Manager

	currentUserGuid string
}

func NewWorker() *Worker {
	dbMgr := db.New()
	return &Worker{dbManager: dbMgr}
}

func (w *Worker) Stop() {
	w.dbManager.Disconnect()
}

func (w *Worker) CompleteRegistration(req *http.Request) error {
	newUser := cmn.User{}
	newUser.Login = req.FormValue("email")
	newUser.Password = req.FormValue("psw")
	newUser.Name = req.FormValue("name")
	newUser.Surname = req.FormValue("lastname")
	newUser.Phone = req.FormValue("phone")

	guid, err := data.Register(newUser, w.dbManager)
	w.currentUserGuid = guid
	return err
}

func (w *Worker) CompleteLogin(req *http.Request) error {
	loggingUser := cmn.User{}
	loggingUser.Login = req.FormValue("email")
	loggingUser.Password = req.FormValue("psw")

	guid, err := data.LoginUser(w.dbManager, loggingUser)
	w.currentUserGuid = guid
	return err
}

func (w *Worker) GalleryHandler(rw http.ResponseWriter, req *http.Request) {
	path := filepath.Join("web", "app", "html", "galleryPage.html")
	tmpl, err := template.ParseFiles(path)

	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	userInfo, err := data.GetInfoAboutUser(w.currentUserGuid, w.dbManager)
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(rw, userInfo)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}
