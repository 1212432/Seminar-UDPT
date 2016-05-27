package main

import (
    "net/http"
    "html/template"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, nil)
}

func main() {
    http.HandleFunc("/index/", viewHandler)
    fs := http.FileServer(http.Dir(""))
	http.Handle("/css/", fs)
	http.Handle("/images/", fs)
    http.ListenAndServe(":8080", nil)
}