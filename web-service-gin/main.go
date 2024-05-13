package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// Struct para representar um usuário
type User struct {
	ID       int
	Username string
	Email    string
}

var connStr = "user=postgres dbname=golang password=123 host=localhost sslmode=disable"

func main() {
	// cria um servidor web já com tratamento de rota
	router := gin.Default()
	// quem vai responder o endpoint /albums é a função getAlbums
	router.GET("/users", getAllUsers)
	router.GET("/albums/:id", getAlbumByID)
	router.DELETE("/albums/:id", deleteAlbumByID)
	router.POST("/user", postAlbums)

	router.Run("localhost:8080")
}

func getAllUsers(c *gin.Context) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT id, username, email FROM users")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	}

	var users []User
	for rows.Next() {
		var user User
		// adiciona no objeto user
		err := rows.Scan(&user.ID, &user.Username, &user.Email)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}
		users = append(users, user) // adiciona no vetor users
	}
	c.IndentedJSON(http.StatusOK, users)
}

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album
	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return // pois tem erro na conversão de JSON para vetor
	}
	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	// não encontrou o álbum
	c.IndentedJSON(http.StatusNotFound,
		gin.H{
			"message": "album not found"})
}

func deleteAlbumByID(c *gin.Context) {
	id := c.Param("id")
	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for i, a := range albums {
		if a.ID == id {
			// Delete the album from the slice.
			albums = append(albums[:i], albums[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "album deleted"})
			return
		}
	}
	// If we got here, it means the album was not found.
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
