package main

import (
	"fmt"
	"github.com/volvinbur1/security-web-app/internal/routing"
	"github.com/volvinbur1/security-web-app/internal/web"
	"log"
	"net/http"
)

var webWorker *web.Worker

func main() {
	webWorker = web.NewWorker()
	router := http.NewServeMux()
	routes(router, webWorker)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}

func routes(mux *http.ServeMux, worker *web.Worker) {
	mux.HandleFunc("/registration", RegistrationHandler)
	mux.HandleFunc("/login", LoginUser)
	mux.HandleFunc("/gallery", worker.GalleryHandler)

	fs := http.FileServer(http.Dir("./web/app/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fs))
}

func RegistrationHandler(rw http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		routing.RegisterPage(rw)
		return
	}

	if req.Method == http.MethodPost {
		err := req.ParseForm()
		if err != nil {
			log.Fatal(err)
		}

		err = webWorker.CompleteRegistration(req)
		if err != nil {
			log.Print(err)
			fmt.Fprintf(rw, err.Error())
		} else {
			http.Redirect(rw, req, "/gallery", 301)
		}
	}
}

func LoginUser(rw http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		routing.LoginPage(rw)
		return
	}

	if req.Method == http.MethodPost {
		err := req.ParseForm()
		if err != nil {
			log.Fatal(err)
		}

		err = webWorker.CompleteLogin(req)
		if err != nil {
			fmt.Fprintf(rw, err.Error())
		} else {
			http.Redirect(rw, req, "/gallery", 301)
		}
	}
}
