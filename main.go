package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album record
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// TODO: save to database
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// get albumns
func getAlbums(ctx *gin.Context) {
	// format json use IndentedJSON, in production use JSON
	ctx.IndentedJSON(http.StatusOK, albums)
}

func createAlbum(ctx *gin.Context) {
	var newAlbum album
	if err := ctx.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
	ctx.IndentedJSON(http.StatusCreated, albums)
}

func getAlbumById(ctx *gin.Context) {
	id := ctx.Param("id")
	for _, a := range albums {
		if a.ID == id {
			ctx.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found with id: " + id})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", createAlbum)
	router.GET("/albums/:id", getAlbumById)
	router.Run(":8080")
}
