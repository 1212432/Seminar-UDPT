package main

import (
    "net/http"
    "html/template"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "contact")
}

func personOneHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "1212432")
}

func personTwoHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "1212439")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
    t, err := template.ParseFiles(tmpl + ".html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    err = t.Execute(w, nil)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func main() {
    http.HandleFunc("/index/", indexHandler)
    http.HandleFunc("/about/", aboutHandler)
    http.HandleFunc("/contact/", contactHandler)
    http.HandleFunc("/1212432/", personOneHandler)
    http.HandleFunc("/1212439/", personTwoHandler)
    fs := http.FileServer(http.Dir(""))
	http.Handle("/css/", fs)
	http.Handle("/images/", fs)
    http.ListenAndServe(":8080", nil)
}