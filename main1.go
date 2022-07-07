package main

import (
    "encoding/json"
    "html/template"
    "net/http"
)

type ContactDetails struct {
    Fname   string `json:"fname"`
    Lname   string `json:"lname"`
}

var details ContactDetails

func main() {
    tmpl := template.Must(template.ParseFiles("Template/index1.html"))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            tmpl.Execute(w, nil)
            return
        }

        details = ContactDetails{
            Fname:   r.FormValue("fname"),
            Lname:   r.FormValue("lname"),
        }

        // form1 := ContactDetails{
        //     Fname: details.Fname,
        //     Lname: details.Lname,
        // }
        // json.NewEncoder(w).Encode(form1)
        
        tmpl.Execute(w, struct{ 
			Success bool 
			Form ContactDetails
			}{true, details})
    })
    http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
        form1 := ContactDetails{
            Fname: details.Fname,
            Lname: details.Lname,
        }

        json.NewEncoder(w).Encode(form1)
    })

    http.ListenAndServe(":8080", nil)
}