package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func index(w http.ResponseWriter, r *http.Request) {
	ck, err := r.Cookie("view-count")
	if err != nil {
		fmt.Println(err)
		http.SetCookie(w, &http.Cookie{
			Name:  "view-count",
			Value: "1",
		})
	} else {
		iValue, _ := strconv.Atoi(ck.Value)
		iValue++
		fmt.Printf("View count:%d\n", iValue)
		ck.Value = fmt.Sprintf("%d", iValue)
		http.SetCookie(w, ck)
	}
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
