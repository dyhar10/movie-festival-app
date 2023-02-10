package main

import (
	"github.com/dyhar10/movie-festival-app/models"
	"github.com/dyhar10/movie-festival-app/routes"
)

func main() {

	db := models.ConfigDB()
	db.AutoMigrate(&models.Movie{})

	r := routes.SetupRoutes(db)
	r.Run()
}
