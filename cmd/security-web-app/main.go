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

	err := http.ListenAndServeTLS(":443", "certificate/localhost.crt", "certificate/localhost.key", router)
	if err != nil {
		log.Fatal(err)
	}
}

func routes(mux *http.ServeMux, worker *web.Worker) {
	mux.HandleFunc("/registration", RegistrationHandler)
	mux.HandleFunc("/login", func(writer http.ResponseWriter, request *http.Request) {
		if worker.LoginHandler(writer, request) {
			http.Redirect(writer, request, "/gallery", 301)
		}
	})
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
