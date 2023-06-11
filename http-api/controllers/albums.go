package controllers

import (
	"github.com/clebersonp/http-api/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

// AlbumGroup is a base url for '/albums' resource
func AlbumGroup(g *gin.RouterGroup) {
	g.GET("", getAlbums)
	g.POST("", postAlbum)
	g.GET("/:id", getAlbumByID)
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	albums := make([]model.Album, 0)
	for _, value := range model.AlbumDb {
		albums = append(albums, value)
	}
	if len(albums) == 0 {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbum adds an album from JSON received in the request body.
func postAlbum(c *gin.Context) {
	var newAlbum model.Album

	// Call BindJSON to bind the received JSON to new Album
	if err := c.BindJSON(&newAlbum); err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Add the new album to the map
	id := uuid.New()
	newAlbum.ID = id.String()
	model.AlbumDb[id.String()] = newAlbum

	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id parameter sent by the client,
// then returns that album as a response
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// looking for an album whose ID value matches the parameter
	album, ok := model.AlbumDb[id]
	if !ok {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}
	c.IndentedJSON(http.StatusOK, album)
}
