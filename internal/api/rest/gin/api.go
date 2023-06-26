package gin

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getID(c *gin.Context) (int, error) {
	// Use context parameters to get the request parameter for the album id
	id := c.Param("id")
	idInt, err := stringToInt(id)
	if err != nil {
		return 0, err
	}

	return idInt, nil
}

// getAlbums route : creates JSON from the slice of album structs, writing the JSON into the response
// `gin.Context` is the most important part of Gin. It carries request details, validates and
// serializes JSON, and more.
func getAlbums(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	// Call Context.IndentedJSON to serialize the struct into JSON and add it to the response.
	// Note that you can replace Context.IndentedJSON with a call to Context.JSON to send more compact
	// JSON.
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var body postAlbumBody

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&body); err != nil {
		return
	}

	// create a new ID to create the new Album
	lastAlbum := albums[len(albums)-1]
	newID := lastAlbum.ID + 1
	newAlbum := album{
		newID, body.Title, body.Artist, body.Price,
	}

	// add the new album to the slice
	addAlbum(newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id, err := getID(c)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	// Loop over the list of albums, looking for an album whose ID value matches the parameter.
	for _, myAlbum := range albums {
		if myAlbum.ID == id {
			c.IndentedJSON(http.StatusOK, myAlbum)
			return
		}
	}

	// otherwise return 404
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album bot found", "id": id})
}

// deleteAlbumByID locates the album whose ID value matches the id in request property
// then deletes the album in backend
func deleteAlbumByID(c *gin.Context) {
	id, err := getID(c)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	for index, myAlbum := range albums {
		if myAlbum.ID == id {
			albums = removeFastByIndex(albums, index)
		}
	}
	c.IndentedJSON(http.StatusOK, albums)
}

// Run defines the API configurations, routes and run the server
func Run() {
	// api init
	// Initialize a Gin router using Default.
	router := gin.Default()

	// CORS config
	// CONFIGURE IT BEFORE ROUTES !
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	// paths declarations
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.DELETE("/albums/:id", deleteAlbumByID)

	// start
	// Use the Run function to attach the router to an http.Server and start the server
	router.Run("localhost:8080")
}
