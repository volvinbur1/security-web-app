package main

import (
	"github.com/volvinbur1/security-web-app/internal/web"
	"log"
	"net/http"
)

func main() {
	webWorker := web.NewWorker()
	router := http.NewServeMux()
	routes(router, webWorker)

	err := http.ListenAndServe(":4040", router)
	if err != nil {
		log.Fatal(err)
	}
}

func routes(mux *http.ServeMux, worker *web.Worker) {
	mux.HandleFunc("/registration", worker.RegistrationHandler)
	mux.HandleFunc("/login", worker.LoginHandler)

	fs := http.FileServer(http.Dir("./web/app/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fs))
}
