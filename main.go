package main

import (
	"net/http"
	"text/template"
)

var tmp1 *template.Template

func init() {
	tmp1 = template.Must(template.ParseFiles("Template/index.html"))
}

type formInfo struct {
	fname string
	lname string
}

func Formhandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		tmp1.Execute(w, nil)
		return
	}
	form := formInfo{
		fname: r.FormValue("fn"),
		lname: r.FormValue("ln"),
	}
	tmp1.Execute(w, struct {
		Success bool
		Form    formInfo
	}{true, form})
}

func main() {
	http.HandleFunc("/", Formhandler)
	http.ListenAndServe(":8080", nil)
}
