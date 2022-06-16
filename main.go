package main

import (
	"log"
	"net/http"
	"text/template"
)

func handler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/register.html")
	if err != nil {
		log.Fatalf("template erro: %v", err)
	}
	if err := t.Execute(w, nil); err != nil {
		log.Fatalf("failed to execute template: %v", err)
	}

}

func main() {
	// TODO: css templateのフォルダを探す機能。あまり理解していないので後で要確認
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources/"))))
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
