package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
)

// album represents data about a record album
type album struct {
	ID 		string 	`json:"id"`
	Title 	string	`json:"title"`
	Artist	string	`json:"artist"`
	Price	float64	`json:"price"`
}

// albums slice to seed record album data
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums responds with the list of all albums as JSON
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// getAlbumById locates the album whose ID value matches the id
// parameter sent by the client, then returns the album as a response
func getAlbumById(c *gin.Context){
	id :=c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID matches the parameter
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// postAlbums adds an album from JSON received in the bequest body
func postAlbums(c *gin.Context){
	var newAlbum album

	// Call BindJSON to bind the received JSON to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func main(){
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.POST("/albums", postAlbums)

	//router.Run("localhost:8080")

	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		eg := v1.Group("/albums")
		{
			eg.GET("/albums", getAlbums)
		}
	}
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	r.Run(":8080")
}

// use 'go run .' in order to start the service
// when getting the dependency, use 'go get .' 

// setting up Swagger
// Step 1: Setup Swagger CLI: go get -u github.com/swaggo/swag/cmd/swag
// 							  go install github.com/swaggo/swag/cmd/swag@latest

// Step 2: Install Gin-Swagger: go get -u github.com/swaggo/gin-swagger
//         Import these packages into the main.go file:
//		   import "github.com/swaggo/gin-swagger"
//		   import "github.com/swaggo/gin-swagger/swaggerFiles"

