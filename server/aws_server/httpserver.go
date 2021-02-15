package main

import (
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

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)
	http.ListenAndServe(":80", nil)
}
