package routes

import (
	"github.com/clebersonp/http-api/controllers"
	"github.com/gin-gonic/gin"
	"log"
)

func HandleRequest() {
	router := gin.Default()
	albumsGroup := router.Group("/albums")
	controllers.AlbumGroup(albumsGroup)

	err := router.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
}
