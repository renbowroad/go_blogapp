package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

func handler(w http.ResponseWriter, r *http.Request) {
	tmp()
	t, err := template.ParseFiles("../templates/login.html")
	if err != nil {
		log.Fatalf("template error: %v", err)
	}
	if err := t.Execute(w, nil); err != nil {
		log.Fatalf("failed to execute template: %v", err)
	}
}

func register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fmt.Println("POST")
	}
	t, err := template.ParseFiles("../templates/register.html")
	if err != nil {
		log.Fatalf("template error: %v", err)
	}
	if err := t.Execute(w, nil); err != nil {
		log.Fatalf("failed to execute template: %v", err)
	}
}

func tmp() {
	db, err := sql.Open("mysql", "kazuhira:password@(172.17.0.2:3306)/practice")
	fmt.Println(db)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
}

func main() {
	// css templateのフォルダを探す機能
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources"))))
	http.HandleFunc("/", handler)
	http.HandleFunc("/register/", register)
	http.ListenAndServe(":8080", nil)
}
