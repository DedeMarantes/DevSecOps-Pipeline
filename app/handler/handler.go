package handler

import (
	"app/db"
	"app/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Aqui vao ter os Cruds
func GetAnimes(c *gin.Context) {
	var animes []model.Anime
	err := db.GetDB().Find(&animes).Error
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve animes"})
		return
	}
	c.JSON(200, animes)
}

func GetAnimeByID(c *gin.Context) {
	var anime model.Anime
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid anime ID"})
		return
	}
	err = db.GetDB().First(&anime, id).Error
	if err != nil {
		c.JSON(404, gin.H{"error": "Anime not found"})
		return
	}
	c.JSON(200, anime)
}

func EditAnime(c *gin.Context) {
	var anime model.Anime
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid anime ID"})
		return
	}
	err = db.GetDB().First(&anime, id).Error
	if err != nil {
		c.JSON(404, gin.H{"error": "Anime not found"})
		return
	}
	if err := c.ShouldBindJSON(&anime); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	anime.ID = id
	err = db.GetDB().Updates(&anime).Error
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to update anime"})
		return
	}
	c.JSON(200, anime)
}

func DeleteAnime(c *gin.Context) {
	var anime model.Anime
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid anime ID"})
		return
	}
	err = db.GetDB().First(&anime, id).Error
	if err != nil {
		c.JSON(404, gin.H{"error": "Anime not found"})
		return
	}
	err = db.GetDB().Delete(&anime).Error
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete anime"})
		return
	}
	c.JSON(200, gin.H{"message": "Anime deleted successfully"})
}

func CreateAnime(c *gin.Context) {
	var anime model.Anime
	if err := c.ShouldBindJSON(&anime); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	err := db.GetDB().Create(&anime).Error
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create anime"})
		return
	}
	c.JSON(201, anime)
}

func HelloWorld(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello, World!"})
}
