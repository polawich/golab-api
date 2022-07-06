package main

import (
	"net/http"
)

var tmp1 * template.Template

type Forms struct{
	fname string
	lname string
}

func init(){
	tmp1 = template.Must(template.ParseFiles("Template/index.html"))
}

func Forms(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		tmp1.Execute(w, nil)
	}
	uforms:=Forms{
		Fname : r.FormValue("fname"),
		Lname : r.FormValue("lname"),
	}
	tmp1.Execute(w,struct {
		Success bool
		Student Forms
	}{true, uforms})
}

func main() {
	http.HandleFunc("/",Forms)
	http.ListenAndServe(":8080",nil)
}