package routing

import (
	"html/template"
	"net/http"
	"path/filepath"
)

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
func RegisterPage(rw http.ResponseWriter) {
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
}
