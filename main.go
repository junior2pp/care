package main

import (
	"net/http"
	"text/template"
	"log"
)

func main() {
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/css/", fs)
	http.Handle("/imgs/", fs)
	http.HandleFunc("/", home)
	http.HandleFunc("/", noticia)
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request)  {
	t,err := template.ParseFiles("./public/html/home.html")
	if err != nil{
		log.Println(err)
	}
	err = t.Execute(w,nil)
	if err != nil{
		log.Println(err)
	}
}

func noticia(w http.ResponseWriter, r *http.Request)  {
	t,err := template.ParseFiles("./public/html/noticias.html")
	if err != nil{
		log.Println(err)
	}
	err = t.Execute(w,nil)
	if err != nil{
		log.Println(err)
	}
}