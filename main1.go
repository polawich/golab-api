// forms.go
package main

import (
    "html/template"
    "net/http"
)

type ContactDetails struct {
    Fname   string
    Lname string
}

func main() {
    tmpl := template.Must(template.ParseFiles("Template/index1.html"))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            tmpl.Execute(w, nil)
            return
        }

        details := ContactDetails{
            Fname:   r.FormValue("fname"),
            Lname: r.FormValue("lname"),
        }

        // // do something with details
        // _ = details

        tmpl.Execute(w, struct{ 
			Success bool 
			Form ContactDetails
			}{true, details})
    })

    http.ListenAndServe(":8080", nil)
}