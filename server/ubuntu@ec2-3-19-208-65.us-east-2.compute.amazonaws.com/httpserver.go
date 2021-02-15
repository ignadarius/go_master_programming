package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

//MyHandler asd
type MyHandler int

func index(w http.ResponseWriter, r *http.Request) {

}

func dog(w http.ResponseWriter, r *http.Request) {
	err := tpl.Execute(w, "Darius")
	if err != nil {
		log.Fatal(err)
	}
}

func me(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Igna Darius")
}

var tpl *template.Template

func init() {
	var err error
	tpl, err = template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)
	http.ListenAndServe(":8080", nil)
}
