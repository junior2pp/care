package main

import "net/http"

func main() {
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/css/", fs)
	http.Handle("/html/", fs)
	http.Handle("/imgs/", fs)
	http.ListenAndServe(":8080", nil)
}