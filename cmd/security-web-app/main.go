package main

import (
	"github.com/volvinbur1/security-web-app/web/pages"
	"log"
	"net/http"
)

func main() {

	router := http.NewServeMux()
	routes(router)

	err := http.ListenAndServe("localhost:4444", router)
	if err != nil {
		log.Fatal(err)
	}
}
func routes(mux *http.ServeMux) {
	mux.HandleFunc("/registration", pages.RegisterPage)
	mux.HandleFunc("/login", pages.LoginPage)
	fs := http.FileServer(http.Dir("./web/app/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	//r.GET("/users", controller.GetUsers)
}
