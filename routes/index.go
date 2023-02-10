package routes

import (
	"github.com/dyhar10/movie-festival-app/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	moviesRg := r.Group("/movies")
	{
		moviesRg.GET("/list", controllers.FindMovies)
		moviesRg.POST("/movies", controllers.CreateMovie)
		moviesRg.GET("/:uuid", controllers.FindMovie)
		moviesRg.PATCH("/:uuid", controllers.UpdateMovie)
	}

	usersRg := r.Group("/users")
	{
		usersRg.GET("/list", controllers.FindMovies)
		usersRg.POST("/movies", controllers.CreateMovie)
		usersRg.GET("/:uuid", controllers.FindMovie)
		usersRg.PATCH("/:uuid", controllers.UpdateMovie)
	}

	return r
}
