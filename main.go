package main

import (
	"example/webservice/controllers"
	"example/webservice/database"

	"github.com/gin-gonic/gin"
)

func logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

func main() {

	router := gin.Default()

	router.Use(logger())

	database.InitDatabase()

	albums := router.Group("/albums")
	{
		albums.GET("", controllers.GetAlbums)

		albums.GET("/:id", controllers.GetAlbumByID)

		albums.POST("", controllers.AddAlbum)

		albums.DELETE("/:id", controllers.DeleteAlbum)

		albums.PUT("/:id", controllers.UpdateAlbum)
	}

	router.Run(":8080")
}
