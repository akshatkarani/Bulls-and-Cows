package main

import (
	"net/http"
	"text/template"
)


func init() {
	tpl = template.Must(template.ParseGlob("templates/*gohtml"))
}


func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/easy/", easy)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":9000", nil)
}
