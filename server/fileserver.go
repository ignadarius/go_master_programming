package main

import (
	"html/template"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {

}

func dogimg(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "dog.jpeg")
}

func dog(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("dog.gohtml")
	if err != nil {
		log.Fatal(err)
	}
	tpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/dog.jpeg/", dogimg)
	http.ListenAndServe(":8080", nil)
}
