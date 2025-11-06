package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	// Parse HTML templates from separate files
	tpl := template.Must(template.ParseFiles("templates/form.html"))
	resultTpl := template.Must(template.ParseFiles("templates/result.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if err := tpl.Execute(w, nil); err != nil {
			log.Println("template error:", err)
		}
	})

	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		if err := r.ParseForm(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("failed to parse form"))
			return
		}

		userInput := r.FormValue("userInput")

		if err := resultTpl.Execute(w, userInput); err != nil {
			log.Println("result template error:", err)
		}
	})

	log.Println("Listening on http://localhost:8080 ...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
