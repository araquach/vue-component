package main

import (
	"github.com/jinzhu/gorm"
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

type Error struct {
	Message string `json:"message"`
}

type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique_index"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
	Token    string `json:"token" gorm:"-"`
}

type TeamMember struct {
	Id    uint   `json:"id" gorm:"primary_key"`
	Slug  string `json:"slug"`
	Name  string `json:"name"`
	Para1 string `json:"para_1"`
	Para2 string `json:"para_2"`
}

type Test struct {
	Id        uint   `json:"id" gorm: "primary_key"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func dbConn() (db *gorm.DB) {
	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	return db
}

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

	db := dbConn()
	db.LogMode(true)
	db.AutoMigrate(&User{}, &Test{}, &TeamMember{})
	db.Close()

	tplIndex = template.Must(template.ParseFiles(
		"views/layouts/main.gohtml",
		"views/pages/index.gohtml"))
	if err != nil {
		panic(err)
	}

	r := muxie.NewMux()
	r.PathCorrection = true
	r.HandleFunc("/*", index)

	// api
	r.HandleFunc("/api/register", apiRegister)
	r.HandleFunc("/api/login", apiLogin)
	r.HandleFunc("/api/team", apiTeam)
	r.HandleFunc("/api/team/:slug", apiTeamMember)
	r.HandleFunc("/api/test", TokenVerifyMiddleWare(apiTest))

	r.Handle("/dist/*file", http.StripPrefix("/dist/", http.FileServer(http.Dir("./dist"))))

	log.Printf("Starting server on %s", port)

	http.ListenAndServe(":"+port, r)
}
