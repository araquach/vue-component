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
		"views/layouts/main.gohtml",
		"views/pages/index.gohtml"))
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
