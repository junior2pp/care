package main

import (
	"net/http"
	"text/template"
	"log"
	"github.com/gorilla/mux"
	"time"
	"fmt"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/home", home)
	r.HandleFunc("/noticia/{id}", noticia)

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("public"))))


	srv := &http.Server{
		Handler:      r,
		Addr:         ":8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	srv.ListenAndServe()
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
type DatosNoticia struct {
	Id int
	Titulo string
	Cuerpo string
	Fecha string
	Autor string
	Correo string
}
func noticia(w http.ResponseWriter, r *http.Request)  {


	vars := mux.Vars(r)
	identificacion := vars["id"]

	db, err:= sql.Open("sqlite3", "./datos.db")
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("SELECT * FROM noticias where id=?")
	if err != nil {
		panic(err)
	}

	rows, err := stmt.Query(identificacion)
	if err != nil {
		panic(err)
	}

	var D DatosNoticia

	for rows.Next() {
		err = rows.Scan(&D.Id, &D.Titulo, &D.Cuerpo, &D.Fecha, &D.Autor, &D.Correo)
		if err != nil {
			panic(err)
		}
	}
	defer rows.Close()
	if D.Id == 0 {
		fmt.Println("Documento no existente")
		t, err := template.ParseFiles("./public/html/noticiaError.html")
		if err != nil{
			log.Println(err)
		}
		err = t.Execute(w, identificacion)
		if err != nil{
			log.Println(err)
		}
	}else {
		t,err := template.ParseFiles("./public/html/noticias.html")
		if err != nil{
			log.Println(err)
		}
		err = t.Execute(w, D)
		if err != nil{
			log.Println(err)
		}
	}


}

func fecha()  {
	año, mes, dia:= time.Now().Date()
	fecha := fmt.Sprintf("%d/%d/%d",año, mes, dia)
	fmt.Println(fecha)
}