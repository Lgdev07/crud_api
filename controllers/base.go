package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Lgdev07/crud_api/middlewares"
	"github.com/Lgdev07/crud_api/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres
)

type App struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (a *App) Initialize(DbHost, DbPort, DbUser, DbName, DbPassword string) {
	var err error
	DBURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)

	a.DB, err = gorm.Open("postgres", DBURI)
	if err != nil {
		fmt.Printf("\n Cannot connect to database %s", DbName)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the database %s", DbName)
	}

	a.DB.Debug().AutoMigrate(
		&models.Store{},
		&models.User{},
	) //database migration

	a.Router = mux.NewRouter().StrictSlash(true)
	a.initializeRoutes()
}

func (a *App) initializeRoutes() {
	a.Router.Use(middlewares.SetContentTypeMiddleware) // setting content-type to json

	a.Router.HandleFunc("/stores", a.createStore).Methods("POST")
	a.Router.HandleFunc("/stores/{id:[0-9]+}", a.updateStore).Methods("PUT")
	a.Router.HandleFunc("/stores/{id:[0-9]+}", a.deleteStore).Methods("DELETE")
	a.Router.HandleFunc("/stores", a.listStore).Methods("GET")
	a.Router.HandleFunc("/stores/{id:[0-9]+}", a.showStore).Methods("GET")
}

func (a *App) RunServer() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	log.Printf("\nServer starting on port %v", port)
	log.Fatal(http.ListenAndServe(":"+port, a.Router))
}
