package main

import (
	"rest-api-crud-gin/src/config"
	"rest-api-crud-gin/src/entitas"
	"rest-api-crud-gin/src/routes"
)

func main() {

	db := config.SetupDB()
	db.AutoMigrate(&entitas.Person{})

	r := routes.SetupRoutes(db)
	r.Run()
}
