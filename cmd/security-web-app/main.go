package main

import (
	"github.com/volvinbur1/security-web-app/internal/routing"
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	routes(router)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}

func routes(mux *http.ServeMux) {
	mux.HandleFunc("/registration", routing.RegisterPage)
	mux.HandleFunc("/login", routing.LoginPage)

	fs := http.FileServer(http.Dir("./web/app/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
}
