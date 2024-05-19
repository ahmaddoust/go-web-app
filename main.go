package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type Hello struct {
	Sale string
	Time string
}

func main() {
	hello := Hello{"Sale Commences Now", time.Now().Format(time.Stamp)}
	outline := template.Must(template.ParseFiles("outline/outline.html"))
	http.Handle("/steady/", http.StripPrefix("/steady/", http.FileServer(http.Dir("steady"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if sale := r.FormValue("sale"); sale != "" {
			hello.Sale = sale
		}
		err := outline.ExecuteTemplate(w, "outline.html", hello)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println(http.ListenAndServe(":8080", nil))
}
