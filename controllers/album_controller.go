package controllers

import (
	"example/webservice/database"
	"example/webservice/models"

	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {
	var albums []models.Album

	database.DB.Find(&albums)

	c.JSON(200, albums)
}
func GetAlbumByID(c *gin.Context) {
	var album models.Album

	// Caso erro ao buscar retorna um 404
	if err := database.DB.First(&album, c.Param("id")).Error; err != nil {
		c.JSON(404, gin.H{"Message": "Album not found"})
		return
	}

	c.JSON(200, album)

}

func AddAlbum(c *gin.Context) {
	var newAlbum models.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	database.DB.Create(&newAlbum)

	c.JSON(201, newAlbum)
}

func DeleteAlbum(c *gin.Context) {
	var album models.Album

	if err := database.DB.First(&album, c.Param("id")).Error; err != nil {
		c.JSON(404, gin.H{"Message": "Album not found"})
		return
	}

	database.DB.Delete(&album, c.Param("id"))

	c.JSON(200, gin.H{"deleted": &album})
}

func UpdateAlbum(c *gin.Context) {
	var a models.Album

	if err := database.DB.First(&a, c.Param("id")).Error; err != nil {
		c.JSON(404, gin.H{"Message": "Album not found"})
		return
	} else if err := c.BindJSON(&a); err != nil {
		c.AbortWithStatus(400)
		return
	}

	database.DB.Save(&a)

	c.JSON(200, a)
}
