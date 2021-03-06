package main

import (
	"os"

	"github.com/Lgdev07/crud_api/controllers"
)

func main() {

	app := controllers.App{}
	app.Initialize(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"))

	app.RunServer()
}
