package controllers

import (
	"example/webservice/models"

	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {
	var albums []models.Album
	models.DB.Find(&albums)

	c.JSON(200, albums)
}
func GetAlbumByID(c *gin.Context) {
	var album models.Album

	// Caso erro ao buscar retorna um 404
	if err := models.DB.First(&album, c.Param("id")).Error; err != nil {
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

	models.DB.Create(&newAlbum)

	c.JSON(201, newAlbum)
}

func DeleteAlbum(c *gin.Context) {
	var album models.Album

	if err := models.DB.First(&album, c.Param("id")).Error; err != nil {
		c.JSON(404, gin.H{"Message": "Album not found"})
		return
	}

	models.DB.Delete(&album, c.Param("id"))

	c.JSON(200, gin.H{"deleted": &album})
}
