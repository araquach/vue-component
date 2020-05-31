package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/kataras/muxie"
	"html/template"
	"log"
	"net/http"
	"os"
)

var (
	tplIndex *template.Template
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	var err error
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	tplIndex = template.Must(template.ParseFiles(
		"views/main.gohtml"))
	if err != nil {
		panic(err)
	}

	r := muxie.NewMux()
	r.PathCorrection = true
	// api
	r.Handle("/dist/*file", http.StripPrefix("/dist/", http.FileServer(http.Dir("./dist"))))

	r.HandleFunc("/*", index)

	log.Printf("Starting server on %s", port)

	http.ListenAndServe(":"+port, r)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := tplIndex.Execute(w, nil); err != nil {
		panic(err)
	}
}
