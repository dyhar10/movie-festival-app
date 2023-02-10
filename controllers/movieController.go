package controllers

import (
	"net/http"

	"github.com/dyhar10/movie-festival-app/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type MovieInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Duration    uint   `json:"duration"`
	Artist      string `json:"artist"`
	Genre       string `json:"genre"`
	URL         string `json:"url"`
}

// GET /movies/list
// Get all list movie
func FindMovies(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var movies []models.Movie
	db.Find(&movies)

	c.JSON(http.StatusOK, gin.H{"data": movies})
}

// POST /movies
// Create new movie data
func CreateMovie(c *gin.Context) {
	// Validate input
	var input MovieInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create movie
	movie := models.Movie{Title: input.Title, Description: input.Description, Duration: input.Duration, Artist: input.Artist, Genre: input.Genre, URL: input.URL}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&movie)

	c.JSON(http.StatusOK, gin.H{"data": movie})
}

// GET /movies/:uuid
// Find a movie by id
func FindMovie(c *gin.Context) {
	var movie models.Movie

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("uuid = ?", c.Param("uuid")).First(&movie).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": movie})
}

// PATCH /movies/:uuid
// Update a movie
func UpdateMovie(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var movie models.Movie
	if err := db.Where("uuid = ?", c.Param("uuid")).First(&movie).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input MovieInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var updatedInput models.Movie

	updatedInput.Title = input.Title
	updatedInput.Description = input.Description
	updatedInput.Duration = input.Duration
	updatedInput.Artist = input.Artist
	updatedInput.Genre = input.Genre
	updatedInput.URL = input.URL

	db.Model(&movie).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": movie})
}
